package registry

import "github.com/fehmicorp/agent/windows/config/types"

func Get() *types.App {
	return instance
}
