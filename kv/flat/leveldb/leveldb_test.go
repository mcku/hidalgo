package leveldb

import (
	"testing"

	"github.com/mcku/hidalgo/kv/flat"
	"github.com/mcku/hidalgo/kv/kvtest"
)

func TestLeveldb(t *testing.T) {
	kvtest.RunTestLocal(t, flat.UpgradeOpenPath(OpenPath), nil)
}
