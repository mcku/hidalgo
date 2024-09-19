//go:build !386 && !arm

package all

// Backends that don't support 32bit

import (
	_ "github.com/mcku/hidalgo/kv/flat/pebble"
)
