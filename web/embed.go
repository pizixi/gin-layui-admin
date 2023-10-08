package web

import "embed"

//go:embed views/*/*.html
var ViewsFS embed.FS

//go:embed static
var StaticFS embed.FS
