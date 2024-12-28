package bella

import (
	"strings"
	"sync/atomic"
)

type node struct {
	key         string
	payload     []byte
	left, right *node
}

// NewNode
func NewNode(key string, payload []byte) *node {
	return &node{
		key:     key,
		payload: payload,
		left:    nil,
		right:   nil,
	}
}

// BST
type BST struct {
	head *node
	size atomic.Int64
}

// New
func New() *BST {
	return &BST{
		head: nil,
		size: atomic.Int64{},
	}
}

// Insert
func (b *BST) Insert(key string, payload []byte) {

	nn := &node{
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

		c := strings.Compare(n.key, nn.key)
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
func (b *BST) InsertNode(nn *node) {

	if b.head == nil {
		b.head = nn
		b.size.Store(1)
		return
	}

	defer b.size.Add(1)

	n := b.head

	for {

		c := strings.Compare(n.key, nn.key)
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
func (b *BST) Search(key string) ([]byte, error) {
	if b.head == nil {
		return nil, ErrNilHead
	}

	n := b.head

	for {

		c := strings.Compare(n.key, key)
		if c == 0 {
			return n.payload, nil
		} else if c == -1 {
			if n.left == nil {
				return nil, ErrNotFound
			}
			n = n.left
		}

		if n.right == nil {
			return nil, ErrNotFound
		}
		n = n.right
	}
}

// InOrderTraversal
func (b *BST) InOrderTraversal() []*node {
	return inOrderTraversal(b.head)
}

func inOrderTraversal(n *node) []*node {
	if n == nil {
		return nil
	}

	if n.left != nil {
		return inOrderTraversal(n.left)
	}

	if n.right != nil {
		return inOrderTraversal(n.right)
	}

	return []*node{n}
}

// Empty
func (b *BST) Empty() {
	b.head = nil
	b.size.Store(0)
}

// Size
func (b *BST) Size() int64 {
	return b.size.Load()
}
