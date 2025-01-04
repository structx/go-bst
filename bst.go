package bst

import (
	"cmp"
	"unsafe"

	"sync/atomic"
)

func zeroGN[T any]() T {
	var t T
	return t
}

type node[K cmp.Ordered, T any] struct {
	key         K
	payload     T
	left, right *node[K, T]
}

// BST
type BtreeGN[K cmp.Ordered, T any] struct {
	head *node[K, T]
	size atomic.Uintptr
}

// Insert
func (b *BtreeGN[K, T]) Insert(key K, payload T) {

	defer b.size.Add(unsafe.Sizeof(payload))

	nn := &node[K, T]{
		key:     key,
		payload: payload,
		left:    nil,
		right:   nil,
	}

	if b.head == nil {
		b.head = nn
		return
	}

	n := b.head

	for {

		if c := cmp.Compare(n.key, nn.key); c == -1 {

			if n.left == nil {
				n.left = nn
				return
			}

			n = n.left
			continue

		} else if c == 0 {
			n.payload = nn.payload
			return
		}

		if n.right == nil {
			n.right = nn
			return
		}

		n = b.head.right

	}
}

// Insert
func (b *BtreeGN[K, T]) InsertNode(nn *node[K, T]) {

	defer b.size.Add(unsafe.Sizeof(nn.payload))

	if b.head == nil {
		b.head = nn
		return
	}

	n := b.head

	for {

		if c := cmp.Compare(n.key, nn.key); c == -1 {

			if n.left == nil {
				n.left = nn
				return
			}

			n = n.left
			continue

		} else if c == 0 {
			n.payload = nn.payload
			return
		}

		if n.right == nil {
			n.right = nn
			return
		}

		n = b.head.right

	}
}

// Search
func (b *BtreeGN[K, T]) Search(key K) (T, error) {
	if b.head == nil {
		return zeroGN[T](), ErrNilHead
	}

	n := b.head

	for {

		if c := cmp.Compare(n.key, key); c == 0 {
			return n.payload, nil

		} else if c == -1 {

			if n.left == nil {
				return zeroGN[T](), ErrNotFound
			}

			n = n.left
			continue
		}

		if n.right == nil {
			return zeroGN[T](), ErrNotFound
		}

		n = n.right
	}
}

// InOrderTraversal
func (b *BtreeGN[K, T]) InOrderTraversal() []*node[K, T] {
	return inOrderTraversal(b.head)
}

func inOrderTraversal[K cmp.Ordered, T any](n *node[K, T]) []*node[K, T] {
	if n == nil {
		return nil
	}

	if n.left != nil {
		return inOrderTraversal(n.left)
	}

	if n.right != nil {
		return inOrderTraversal(n.right)
	}

	return []*node[K, T]{n}
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
