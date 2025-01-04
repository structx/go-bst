package bella_test

import (
	"testing"

	"github.com/structx/bella"
)

var (
	btree *bella.BtreeGN[int, []byte]
)

func init() {
	btree = &bella.BtreeGN[int, []byte]{}
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		btree.Insert(i, []byte("helloworld"))
	}
}
