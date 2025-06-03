//go:build required

package unarr

import (
	_ "github.com/Servicewall/go-unarr/unarrc/external"
	_ "github.com/Servicewall/go-unarr/unarrc/external/bzip2"
	_ "github.com/Servicewall/go-unarr/unarrc/external/unarr"
	_ "github.com/Servicewall/go-unarr/unarrc/external/unarr/_7z"
	_ "github.com/Servicewall/go-unarr/unarrc/external/unarr/common"
	_ "github.com/Servicewall/go-unarr/unarrc/external/unarr/lzmasdk"
	_ "github.com/Servicewall/go-unarr/unarrc/external/unarr/rar"
	_ "github.com/Servicewall/go-unarr/unarrc/external/unarr/tar"
	_ "github.com/Servicewall/go-unarr/unarrc/external/unarr/zip"
	_ "github.com/Servicewall/go-unarr/unarrc/external/zlib"
)
