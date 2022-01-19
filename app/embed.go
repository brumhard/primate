package app

import "embed"

//go:embed build/web
var Static embed.FS
