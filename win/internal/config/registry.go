package config

import (
	"context"
	"sync"
)

var (
	// Store contexts for multiple windows/apps
	contexts = make(map[int]context.Context)
	// Use RWMutex so we can support RLock/RUnlock
	mu sync.RWMutex
)

// SetContext saves or updates the context for a specific ID
func SetContext(id int, ctx context.Context) {
	mu.Lock()
	defer mu.Unlock()
	contexts[id] = ctx
}

// GetContext retrieves the context for a specific ID
func GetContext(id int) context.Context {
	mu.RLock()
	defer mu.RUnlock()
	return contexts[id]
}

// GetAllContexts returns a slice of all active contexts
func GetAllContexts() []context.Context {
	mu.RLock()
	defer mu.RUnlock()

	list := make([]context.Context, 0, len(contexts))
	for _, ctx := range contexts {
		list = append(list, ctx)
	}
	return list
}

// DeleteContext removes a context from the map (useful when closing a window)
func DeleteContext(id int) {
	mu.Lock()
	defer mu.Unlock()
	delete(contexts, id)
}
