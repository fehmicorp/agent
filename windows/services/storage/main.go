package storage

// Store defines the standard CRUD interface for key-value configurations
type Store interface {
	Init() error
	Set(key, value string) error
	Get(key string) (string, error)
	Delete(key string) error
	Close() error
}
