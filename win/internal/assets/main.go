package assets

import (
	"embed"
	_ "embed"
)

//go:embed fav.ico
var Favicon []byte

//go:embde dashboard/*
var Dashboard embed.FS
