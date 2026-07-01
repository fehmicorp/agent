package bbolt

import (
	"fmt"
	"path/filepath"

	"github.com/fehmicorp/agent/windows/storage/fs"
	goBolt "go.etcd.io/bbolt"
)

type BoltStore struct {
	dbName string
	db     *goBolt.DB
	bucket []byte
}

func NewBoltStore(dbName string, bucketName string) *BoltStore {
	return &BoltStore{
		dbName: dbName,
		bucket: []byte(bucketName),
	}
}

// Init ensures directories exist, mounts the database file, and pre-provisions the targeted bucket setup
func (b *BoltStore) Init(dir string) (bool, error) {
	_, err := fs.EnsureDir(dir)
	if err != nil {
		return false, fmt.Errorf("failed to provision database directory path via fs module: %w", err)
	}
	b.db, err = goBolt.Open(filepath.Join(dir, b.dbName), 0600, nil)
	if err != nil {
		return false, err
	}

	// Create bucket atomically inside initialization transactions scope
	err = b.db.Update(func(tx *goBolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists(b.bucket)
		return err
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

// Set inserts or updates a key value record cleanly (Create/Update)
func (b *BoltStore) Set(key, value string) error {
	if b.db == nil {
		return fmt.Errorf("bbolt engine instance not initialized")
	}
	return b.db.Update(func(tx *goBolt.Tx) error {
		bkt := tx.Bucket(b.bucket)
		if bkt == nil {
			return fmt.Errorf("configured bucket not found")
		}
		return bkt.Put([]byte(key), []byte(value))
	})
}

// Get pulls value arrays matching the target key string (Read)
func (b *BoltStore) Get(key string) (string, error) {
	if b.db == nil {
		return "", fmt.Errorf("bbolt engine instance not initialized")
	}
	var val []byte
	err := b.db.View(func(tx *goBolt.Tx) error {
		bkt := tx.Bucket(b.bucket)
		if bkt == nil {
			return fmt.Errorf("configured bucket not found")
		}
		val = bkt.Get([]byte(key))
		return nil
	})
	return string(val), err
}

// Delete completely drops a key pair out of active indexes sheets (Delete)
func (b *BoltStore) Delete(key string) error {
	if b.db == nil {
		return fmt.Errorf("bbolt engine instance not initialized")
	}
	return b.db.Update(func(tx *goBolt.Tx) error {
		bkt := tx.Bucket(b.bucket)
		if bkt == nil {
			return fmt.Errorf("configured bucket not found")
		}
		return bkt.Delete([]byte(key))
	})
}

// Close flushes background system commits and frees platform locks cleanly
func (b *BoltStore) Close() error {
	if b.db != nil {
		return b.db.Close()
	}
	return nil
}

// Push Appends a log item to the tail end of your queue structure
func (b *BoltStore) Push(value string) error {
	return b.db.Update(func(tx *goBolt.Tx) error {
		bkt := tx.Bucket(b.bucket)
		// Get next sequential ID (Auto-incrementing)
		id, _ := bkt.NextSequence()

		// Convert uint64 sequence ID to a fixed-width string byte key for ordering
		key := fmt.Sprintf("%020d", id)
		return bkt.Put([]byte(key), []byte(value))
	})
}

// Pop extracts the oldest entry off the head of the queue pipeline
func (b *BoltStore) Pop() (string, error) {
	var value string
	err := b.db.Update(func(tx *goBolt.Tx) error {
		bkt := tx.Bucket(b.bucket)
		cursor := bkt.Cursor()

		// Move to absolute oldest entry at start of bucket index layout
		k, v := cursor.First()
		if k == nil {
			return nil // Queue is empty
		}

		value = string(v)
		return cursor.Delete() // Delete item immediately to complete pop tracking cycle
	})
	return value, err
}
