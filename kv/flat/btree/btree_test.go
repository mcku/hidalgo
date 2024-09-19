package btree

import (
	"testing"

	"github.com/mcku/hidalgo/kv"
	"github.com/mcku/hidalgo/kv/flat"
	"github.com/mcku/hidalgo/kv/kvtest"
)

func TestBtree(t *testing.T) {
	kvtest.RunTest(t, func(t testing.TB) kv.KV {
		return flat.Upgrade(New())
	}, &kvtest.Options{
		NoLocks: true,
		NoTx:    true,
	})
}
