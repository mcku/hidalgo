package badger

import (
	"testing"

	"github.com/mcku/hidalgo/kv/flat"
	"github.com/mcku/hidalgo/kv/kvtest"
)

func TestBadger(t *testing.T) {
	kvtest.RunTestLocal(t, flat.UpgradeOpenPath(OpenPath), nil)
}
