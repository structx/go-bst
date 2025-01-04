package bst

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)

// ToFile
func (b *BtreeGN[K, T]) ToFile(filePath string) error {

	fp := filepath.Clean(filePath)
	f, err := os.OpenFile(fp, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0600)
	if err != nil {
		return fmt.Errorf("os.OpenFile: %v", err)
	}
	defer func() { _ = f.Close() }()

	w := bufio.NewWriter(f)
	defer func() { _ = w.Flush() }()

	inOrderTraversal(b.head, func(n *Node[K, T]) error {

		var left, right *K = nil, nil
		if n.left != nil {
			left = &n.left.Key
		}

		if n.right != nil {
			right = &n.right.Key
		}

		entry := &exportEntry[K, T]{
			Key:     n.Key,
			Payload: n.Payload,
			Left:    left,
			Right:   right,
		}

		jsonbytes, err := json.Marshal(entry)
		if err != nil {
			return fmt.Errorf("json.Marshal: %v", err)
		}

		jsonbytes = append(jsonbytes, '\n')

		_, err = w.Write(jsonbytes)
		if err != nil {
			return fmt.Errorf("unable to write entry: %v", err)
		}

		return nil
	})

	return nil
}
