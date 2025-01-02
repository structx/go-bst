package bella

import (
	"cmp"
	"sync/atomic"
)

type node[K cmp.Ordered, T any] struct {
	key         K
	payload     T
	left, right *node[K, T]
}

// BST
type BST[K cmp.Ordered, T any] struct {
	head *node[K, T]
	size atomic.Int64
}

// Insert
func (b *BST[K, T]) Insert(key K, payload T) {

	nn := &node[K, T]{
		key:     key,
		payload: payload,
		left:    nil,
		right:   nil,
	}

	if b.head == nil {
		b.head = nn
		b.size.Store(1)
		return
	}

	defer b.size.Add(1)

	n := b.head

	for {

		c := compare(n.key, nn.key)
		if c == -1 {

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
func (b *BST[K, T]) InsertNode(nn *node[K, T]) {

	if b.head == nil {
		b.head = nn
		b.size.Store(1)
		return
	}

	defer b.size.Add(1)

	n := b.head

	for {

		c := compare(n.key, nn.key)
		if c == -1 {
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
func (b *BST[K, T]) Search(key K) (T, error) {
	if b.head == nil {
		return zero[T](), ErrNilHead
	}

	n := b.head

	for {

		c := compare(n.key, key)
		if n.key == key {
			return n.payload, nil
		} else if c == -1 {
			if n.left == nil {
				return zero[T](), ErrNotFound
			}
			n = n.left
		}

		if n.right == nil {
			return zero[T](), ErrNotFound
		}
		n = n.right
	}
}

// InOrderTraversal
func (b *BST[K, T]) InOrderTraversal() []*node[K, T] {
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

// Empty
func (b *BST[K, T]) Empty() {
	b.head = nil
	b.size.Store(0)
}

// Size
func (b *BST[K, T]) Size() int64 {
	return b.size.Load()
}
