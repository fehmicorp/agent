package main

import "sync"

var (
	mu sync.RWMutex
)

func AppTitle(id int) string {
	if fn, ok := Get(id); ok {
		return fn.Title
	}
	return ""
}

func Get(id int) (config.Functions, bool) {
	mu.RLock()
	defer mu.RUnlock()
	fn, ok := registry[id]
	return fn, ok
}

func AppTitle(id int) string {

	if fn, ok := Get(id); ok {
		return fn.Title
	}

	return ""
}

func AppLayout(id int) config.Layout {
	if fn, ok := Get(id); ok {
		return fn.Layout
	}
	return config.Layout{}
}

func AppHidden(id int) bool {

	if fn, ok := Get(id); ok {
		return fn.Hidden
	}

	return false
}
