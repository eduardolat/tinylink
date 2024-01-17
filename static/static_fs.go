package static

import "embed"

//go:embed *
var StaticFs embed.FS

//go:embed favicon.ico
var Favicon []byte
