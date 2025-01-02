package bella_test

import (
	"testing"

	"github.com/structx/bella"
)

var (
	btree *bella.BST[int, []byte]
)

func init() {
	btree = &bella.BST[int, []byte]{}
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		btree.Insert(i, []byte("helloworld"))
	}
}
