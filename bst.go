package bst

import (
	"cmp"
	"sync/atomic"
	"unsafe"
)

func zeroGN[T any]() T {
	var t T
	return t
}

// Node
type Node[K cmp.Ordered, T any] struct {
	Key         K
	Payload     T
	left, right *Node[K, T]
}

// BST
type BtreeGN[K cmp.Ordered, T any] struct {
	head *Node[K, T]
	size atomic.Uintptr
}

// Insert
func (b *BtreeGN[K, T]) Insert(key K, payload T) *Node[K, T] {
	defer b.size.Add(unsafe.Sizeof(payload))

	nn := &Node[K, T]{
		Key:     key,
		Payload: payload,
		left:    nil,
		right:   nil,
	}

	if b.head == nil {
		b.head = nn
		return nn
	}

	return insertInOrder(b.head, nn)
}

func insertInOrder[K cmp.Ordered, T any](n, nn *Node[K, T]) *Node[K, T] {

	if n.Key == nn.Key {
		n = nn
		return nn
	}

	if cmp.Less(nn.Key, n.Key) {
		if n.left == nil {
			n.left = nn
			return nn
		}
		return insertInOrder(n.left, nn)
	}

	if n.right == nil {
		n.right = nn
		return nn
	}

	return insertInOrder(n.right, nn)
}

// Search
func (b *BtreeGN[K, T]) Search(key K) (T, error) {
	return searchInOrder(b.head, key)
}

func searchInOrder[K cmp.Ordered, T any](n *Node[K, T], key K) (T, error) {
	if n == nil {
		return zeroGN[T](), ErrNotFound
	}

	if n.Key == key {
		return n.Payload, nil
	}

	if cmp.Less(key, n.Key) {
		return searchInOrder(n.left, key)
	}

	return searchInOrder(n.right, key)
}

// InOrderTraversal
func (b *BtreeGN[K, T]) InOrderTraversal(f func(*Node[K, T]) error) {
	inOrderTraversal(b.head, f)
}

func inOrderTraversal[K cmp.Ordered, T any](n *Node[K, T], f func(*Node[K, T]) error) {
	if n == nil {
		return
	}

	inOrderTraversal(n.left, f)
	_ = f(n)
	inOrderTraversal(n.right, f)
}

// Flush
func (b *BtreeGN[K, T]) Flush() {
	b.head = nil
	b.size.Store(0)
}

// Size
func (b *BtreeGN[K, T]) Size() uintptr {
	return b.size.Load()
}
