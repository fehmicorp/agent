package sqlite

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	_ "modernc.org/sqlite" // Pure Go SQLite driver (No CGO required)
)

type SQLiteStore struct {
	dbName string
	db     *sql.DB
	table  string
}

type TblQuery struct {
	Key        string `json:"key"`        // Column Name (e.g., "id", "device_token")
	Type       string `json:"type"`       // Data Type (e.g., "TEXT", "INTEGER")
	Preference string `json:"preference"` // Constraints (e.g., "PRIMARY KEY", "NOT NULL")
}

// NewSQLiteStore creates a new storage instance with customizable file names and table names
func NewSQLiteStore(dbName string, tableName string) *SQLiteStore {
	return &SQLiteStore{
		dbName: dbName,
		table:  tableName,
	}
}

// Init ensures directories exist, opens the SQLite database file, and configures the schema dynamically
func (s *SQLiteStore) Init(dir string, columns []TblQuery) (bool, error) {
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

	// Dynamically assemble the columns query list
	var colDefinitions []string
	for _, col := range columns {
		def := fmt.Sprintf("%s %s", col.Key, col.Type)
		if col.Preference != "" {
			def = fmt.Sprintf("%s %s", def, col.Preference)
		}
		colDefinitions = append(colDefinitions, def)
	}

	// Always append an optional audit column if not already declared
	colDefinitions = append(colDefinitions, "updated_at DATETIME DEFAULT CURRENT_TIMESTAMP")

	query := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (%s);", s.table, strings.Join(colDefinitions, ", "))

	_, err = s.db.Exec(query)
	if err != nil {
		_ = s.db.Close()
		return false, fmt.Errorf("failed to initialize dynamic schema constraints: %w", err)
	}

	return true, nil
}

// Set handles the default key-value schema fallback setup (backward compatibility)
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

// Get queries a string parameter matching the target key context rows
func (s *SQLiteStore) Get(key string) (string, error) {
	if s.db == nil {
		return "", fmt.Errorf("sqlite database instance pool is not initialized")
	}

	query := fmt.Sprintf(`SELECT value FROM %s WHERE key = ? LIMIT 1;`, s.table)
	var value string

	err := s.db.QueryRow(query, key).Scan(&value)
	if err == sql.ErrNoRows {
		return "", nil
	}
	if err != nil {
		return "", err
	}

	return value, nil
}

// Delete completely drops a key row mapping record out of active tables sheets
func (s *SQLiteStore) Delete(key string) error {
	if s.db == nil {
		return fmt.Errorf("sqlite database instance pool is not initialized")
	}

	query := fmt.Sprintf(`DELETE FROM %s WHERE key = ?;`, s.table)
	_, err := s.db.Exec(query, key)
	return err
}

// ============================================================================
// Dynamic CRUD Engine Operations (For any custom schemas)
// ============================================================================

// InsertDynamic takes an arbitrary map of columns and values to construct a secure dynamic insert statement
func (s *SQLiteStore) InsertDynamic(data map[string]any) error {
	if s.db == nil {
		return fmt.Errorf("sqlite database instance pool is not initialized")
	}
	if len(data) == 0 {
		return fmt.Errorf("cannot insert empty records map")
	}

	var columns []string
	var placeholders []string
	var args []any

	for col, val := range data {
		columns = append(columns, col)
		placeholders = append(placeholders, "?")
		args = append(args, val)
	}

	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);",
		s.table,
		strings.Join(columns, ", "),
		strings.Join(placeholders, ", "),
	)

	_, err := s.db.Exec(query, args...)
	return err
}

// UpdateDynamic updates generic structures tracking parameters using custom conditional where clauses
func (s *SQLiteStore) UpdateDynamic(data map[string]any, whereClause string, whereArgs ...any) error {
	if s.db == nil {
		return fmt.Errorf("sqlite database instance pool is not initialized")
	}
	if len(data) == 0 {
		return fmt.Errorf("cannot update empty data metrics mapping")
	}

	var setStatements []string
	var args []any

	for col, val := range data {
		setStatements = append(setStatements, fmt.Sprintf("%s = ?", col))
		args = append(args, val)
	}

	args = append(args, whereArgs...)

	query := fmt.Sprintf("UPDATE %s SET %s", s.table, strings.Join(setStatements, ", "))
	if whereClause != "" {
		query = fmt.Sprintf("%s WHERE %s;", query, whereClause)
	}

	_, err := s.db.Exec(query, args...)
	return err
}

// DeleteDynamic completely clears data rows filtered dynamically by conditions
func (s *SQLiteStore) DeleteDynamic(whereClause string, args ...any) error {
	if s.db == nil {
		return fmt.Errorf("sqlite database instance pool is not initialized")
	}

	query := fmt.Sprintf("DELETE FROM %s", s.table)
	if whereClause != "" {
		query = fmt.Sprintf("%s WHERE %s;", query, whereClause)
	}

	_, err := s.db.Exec(query, args...)
	return err
}

// Close flushes open database connections safely
func (s *SQLiteStore) Close() error {
	if s.db != nil {
		return s.db.Close()
	}
	return nil
}
