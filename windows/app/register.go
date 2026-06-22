package main

import (
	"context"
	"embed"
	"sync"

	"github.com/fehmicorp/agent/windows/config/types"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	mu     sync.RWMutex
	reg    = make(map[int]types.Window)
	assets = make(map[int]embed.FS)
)

func Register(fn types.Window) {
	mu.Lock()
	defer mu.Unlock()

	reg[fn.Id] = fn
}

func RegisterAssets(id int, fs embed.FS) {
	mu.Lock()
	defer mu.Unlock()

	assets[id] = fs
}

func RegisterMany(items []types.Window) {
	mu.Lock()
	defer mu.Unlock()

	for _, item := range items {
		reg[item.Id] = item
	}
}
func Get(id int) (types.Window, bool) {
	mu.RLock()
	defer mu.RUnlock()
	fn, ok := reg[id]
	return fn, ok
}

func AppTitle(id int) string {

	if fn, ok := Get(id); ok {
		return fn.Title
	}

	return ""
}

func AppLayout(id int) types.Layout {
	if fn, ok := Get(id); ok {
		return fn.Layout
	}
	return types.Layout{}
}

func AppHidden(id int) bool {

	if fn, ok := Get(id); ok {
		return fn.Hidden
	}

	return false
}

func AppOnStartup(id int) bool {

	if fn, ok := Get(id); ok {
		return fn.Startup
	}

	return false
}

func AppOnBeforeClose(id int) bool {

	if fn, ok := Get(id); ok {
		return fn.BeforeClose
	}

	return false
}

func AppRoute(id int) string {
	if fn, ok := Get(id); ok {
		return fn.Route
	}
	return "/"
}

func ShowRoute(
	ctx context.Context,
	route string,
) {
	if ctx == nil {
		return
	}
	runtime.WindowShow(ctx)
	runtime.WindowUnminimise(ctx)
	runtime.WindowExecJS(
		ctx,
		`window.location.hash="`+route+`";`,
	)
}
