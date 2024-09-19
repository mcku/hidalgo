package tuplekv_test

import (
	"testing"

	"github.com/mcku/hidalgo/kv/flat"
	"github.com/mcku/hidalgo/kv/flat/btree"
	"github.com/mcku/hidalgo/tuple"
	tuplekv "github.com/mcku/hidalgo/tuple/kv"
	"github.com/mcku/hidalgo/tuple/tupletest"
)

func TestKV2Tuple(t *testing.T) {
	tupletest.RunTest(t, func(t testing.TB) tuple.Store {
		kdb := btree.New()
		db := tuplekv.New(flat.Upgrade(kdb))
		return db
	}, &tupletest.Options{
		NoLocks: true,
	})
}
