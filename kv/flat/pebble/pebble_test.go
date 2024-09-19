//go:build !386 && !arm

package pebble

import (
	"testing"

	"github.com/mcku/hidalgo/kv/flat"
	"github.com/mcku/hidalgo/kv/kvtest"
)

func TestPebble(t *testing.T) {
	kvtest.RunTestLocal(t, flat.UpgradeOpenPath(OpenPath), &kvtest.Options{
		NoTx: true,
	})
}
