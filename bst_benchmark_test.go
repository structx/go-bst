package bella_test

import (
	"strconv"
	"testing"

	"github.com/structx/bella"
)

var (
	btree *bella.BST
)

func init() {
	btree = bella.New()
}

func BenchmarkInsert(b *testing.B) {
	for i := 0; i < b.N; i++ {
		k := strconv.Itoa(i)
		btree.Insert(k, []byte(k))
	}
}
