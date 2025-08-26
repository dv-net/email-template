package assets

import (
	"embed"
)

//go:embed all:mustache all:i18n
var Assets embed.FS
