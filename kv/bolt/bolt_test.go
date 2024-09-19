package bolt

import (
	"path/filepath"
	"testing"

	"github.com/mcku/hidalgo/kv"
	"github.com/mcku/hidalgo/kv/kvtest"
)

func TestBolt(t *testing.T) {
	kvtest.RunTestLocal(t, func(path string) (kv.KV, error) {
		path = filepath.Join(path, "bolt.db")
		return OpenPath(path)
	}, nil)
}
