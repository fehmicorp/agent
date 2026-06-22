package registry

import (
	"context"
	"sync"
)

var (
	contexts = make(map[int]context.Context)
	mu       sync.RWMutex
)

// SetContext stores a context by ID.
func SetContext(id int, ctx context.Context) {
	mu.Lock()
	defer mu.Unlock()
	contexts[id] = ctx
}

// GetContext retrieves a context by ID.
func GetContext(id int) context.Context {
	mu.RLock()
	defer mu.RUnlock()
	return contexts[id]
}

// GetAllContexts returns a slice of all stored contexts.
func GetAllContexts() []context.Context {
	mu.RLock()
	defer mu.RUnlock()
	list := make([]context.Context, 0, len(contexts))
	for _, ctx := range contexts {
		list = append(list, ctx)
	}
	return list
}

// DeleteContext removes a context by ID.
func DeleteContext(id int) {
	mu.Lock()
	defer mu.Unlock()
	delete(contexts, id)
}
