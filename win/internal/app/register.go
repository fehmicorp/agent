package app

import (
	"context"
	"embed"
	"sync"
	"win/internal/config"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

var (
	mu       sync.RWMutex
	registry = make(map[int]config.Functions)
	assets   = make(map[int]embed.FS)
)

func Register(fn config.Functions) {
	mu.Lock()
	defer mu.Unlock()

	registry[fn.Id] = fn
}

func RegisterAssets(id int, fs embed.FS) {
	mu.Lock()
	defer mu.Unlock()

	assets[id] = fs
}

func RegisterMany(items []config.Functions) {
	mu.Lock()
	defer mu.Unlock()

	for _, item := range items {
		registry[item.Id] = item
	}
}

func Get(id int) (config.Functions, bool) {

	mu.RLock()
	defer mu.RUnlock()

	fn, ok := registry[id]

	return fn, ok
}

func MustGet(id int) config.Functions {

	fn, ok := Get(id)

	if !ok {
		panic("app function not registered")
	}

	return fn
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

func AppAssets(id int) embed.FS {

	mu.RLock()
	defer mu.RUnlock()

	if fs, ok := assets[id]; ok {
		return fs
	}

	return embed.FS{}
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
