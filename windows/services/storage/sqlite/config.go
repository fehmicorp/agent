package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite" // Pure Go SQLite driver (No CGO required)
)

type SQLiteStore struct {
	dbName string
	db     *sql.DB
	table  string
}

type TblQuery struct {
	Key        string `json:"key"`
	Type       string `json:"type"`
	Preference string `json:"preference"`
}

// NewSQLiteStore creates a new storage instance with customizable file names and table names
func NewSQLiteStore(dbName string, tableName string) *SQLiteStore {
	return &SQLiteStore{
		dbName: dbName,
		table:  tableName,
	}
}

// Init ensures directories exist, opens the SQLite database file, and configures the schema
func (s *SQLiteStore) Init(dir string, tblQuery TblQuery) (bool, error) {
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		return false, fmt.Errorf("failed to provision database folder path: %w", err)
	}

	dbPath := filepath.Join(dir, s.dbName)
	var err error

	// Open connection to SQLite database file
	s.db, err = sql.Open("sqlite", dbPath)
	if err != nil {
		return false, fmt.Errorf("failed to open sqlite database file: %w", err)
	}
	_, err = s.db.Exec(query)
	if err != nil {
		_ = s.db.Close()
		return false, fmt.Errorf("failed to initialize schema constraints: %w", err)
	}

	return true, nil
}

// Set inserts a record or updates its content seamlessly if the key already exists (Create/Update)
func (s *SQLiteStore) Set(key, value string) error {
	if s.db == nil {
		return fmt.Errorf("sqlite database instance pool is not initialized")
	}

	query := fmt.Sprintf(`
	INSERT INTO %s (key, value, updated_at) 
	VALUES (?, ?, CURRENT_TIMESTAMP)
	ON CONFLICT(key) DO UPDATE SET 
		value = excluded.value,
		updated_at = CURRENT_TIMESTAMP;`, s.table)

	_, err := s.db.Exec(query, key, value)
	return err
}

// Get queries a string parameter matching the target key context rows (Read)
func (s *SQLiteStore) Get(key string) (string, error) {
	if s.db == nil {
		return "", fmt.Errorf("sqlite database instance pool is not initialized")
	}

	query := fmt.Sprintf(`SELECT value FROM %s WHERE key = ? LIMIT 1;`, s.table)
	var value string

	err := s.db.QueryRow(query, key).Scan(&value)
	if err == sql.ErrNoRows {
		return "", nil // Target key missing, safely return an empty string to keep driver consistency
	}
	if err != nil {
		return "", err
	}

	return value, nil
}

// Delete completely drops a key row mapping record out of active tables sheets (Delete)
func (s *SQLiteStore) Delete(key string) error {
	if s.db == nil {
		return fmt.Errorf("sqlite database instance pool is not initialized")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE key = ?;`, s.table)
	_, err := s.db.Exec(query, key)
	return err
}

// Close flushes open database connections and cleans up background thread resources safely
func (s *SQLiteStore) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
