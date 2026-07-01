package buntdb

import (
	"fmt"
	"time"

	"github.com/tidwall/buntdb"
)

type MemoryStore struct {
	db *buntdb.DB
}

// NewMemoryStore initializes a pure, lightning-fast in-memory BuntDB instance
func NewMemoryStore() (*MemoryStore, error) {
	// Passing ":memory:" ensures zero disk I/O, operating entirely within RAM
	db, err := buntdb.Open(":memory:")
	if err != nil {
		return nil, fmt.Errorf("failed to open in-memory buntdb: %w", err)
	}

	store := &MemoryStore{db: db}

	// Pre-create a spatial/custom index for queue order tracking based on custom sequences
	err = db.Update(func(tx *buntdb.Tx) error {
		// Create a numeric index named "queue_order" on keys matching "queue:*"
		return tx.CreateIndex("queue_order", "queue:*", buntdb.IndexString)
	})
	if err != nil {
		return nil, fmt.Errorf("failed to initialize queue index: %w", err)
	}

	return store, nil
}

// ============================================================================
// Core CRUD Operations
// ============================================================================

// Set saves a key-value pair into RAM permanently until application closes (Create/Update)
func (m *MemoryStore) Set(key, value string) error {
	return m.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, nil)
		return err
	})
}

// SetWithTTL saves a key-value item that self-deletes automatically after duration expiry
func (m *MemoryStore) SetWithTTL(key, value string, duration time.Duration) error {
	return m.db.Update(func(tx *buntdb.Tx) error {
		_, _, err := tx.Set(key, value, &buntdb.SetOptions{Expires: true, TTL: duration})
		return err
	})
}

// Get reads a value array matching the targeted string key (Read)
func (m *MemoryStore) Get(key string) (string, error) {
	var val string
	err := m.db.View(func(tx *buntdb.Tx) error {
		var err error
		val, err = tx.Get(key)
		return err
	})
	if err == buntdb.ErrNotFound {
		return "", nil // Key missing, gracefully return empty string without breaking execution
	}
	return val, err
}

// Delete permanently drops a record from memory tables (Delete)
func (m *MemoryStore) Delete(key string) error {
	return m.db.Update(func(tx *buntdb.Tx) error {
		_, err := tx.Delete(key)
		if err == buntdb.ErrNotFound {
			return nil // Already absent
		}
		return err
	})
}

// ============================================================================
// Advanced Queue Subsystem (Redis-like LPUSH / RPOP pattern)
// ============================================================================

// PushQueue appends an item to the tail end of an ephemeral memory queue pipeline
func (m *MemoryStore) PushQueue(queueName, value string, ttl time.Duration) error {
	return m.db.Update(func(tx *buntdb.Tx) error {
		// Use a precise Unix timestamp prefix to keep strict chronological order in the index
		seq := time.Now().UnixNano()
		key := fmt.Sprintf("queue:%s:%d", queueName, seq)

		var opts *buntdb.SetOptions
		if ttl > 0 {
			opts = &buntdb.SetOptions{Expires: true, TTL: ttl}
		}

		_, _, err := tx.Set(key, value, opts)
		return err
	})
}

// PopQueue pulls and eliminates the oldest available entry off the queue (FIFO ordering)
func (m *MemoryStore) PopQueue(queueName string) (string, error) {
	var targetKey, targetValue string

	err := m.db.Update(func(tx *buntdb.Tx) error {
		// Query our predefined index to find items inside this queue namespace
		var err error
		tx.Ascend("queue_order", func(key, value string) bool {
			if fmt.Sprintf("queue:%s:", queueName) == key[:len(queueName)+7] {
				targetKey = key
				targetValue = value
				return false // Break iteration, found the oldest item
			}
			return true // Keep searching matching patterns
		})

		// If a matching item is found, delete it from memory immediately to finish the "Pop" action
		if targetKey != "" {
			_, err = tx.Delete(targetKey)
			return err
		}
		return nil
	})

	return targetValue, err
}

// Close destroys the database handler cleanly
func (m *MemoryStore) Close() error {
	if m.db != nil {
		return m.db.Close()
	}
	return nil
}
