package bst_test

import (
	"testing"

	"github.com/structx/go-bst"
)

var (
	btree *bst.BtreeGN[int, []byte]
)

func init() {
	btree = &bst.BtreeGN[int, []byte]{}
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		btree.Insert(i, []byte("helloworld"))
	}
}
