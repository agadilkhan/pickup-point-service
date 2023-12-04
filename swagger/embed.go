package swagger

import (
	"embed"
	_ "embed"
)

//go:embed OpenAPI/*
var OpenAPI embed.FS
