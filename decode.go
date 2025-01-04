package bst

import (
	"bufio"
	"bytes"
	"cmp"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type exportEntry[K cmp.Ordered, T any] struct {
	Key     K  `json:"key"`
	Payload T  `json:"payload"`
	Left    *K `json:"left,omitempty"`
	Right   *K `json:"right,omitempty"`
}

// FromFile
func FromFile[K cmp.Ordered, T any](filePath string) (*BtreeGN[K, T], error) {

	f, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return nil, fmt.Errorf("os.Open: %v", err)
	}
	defer func() { _ = f.Close() }()

	var (
		scanner = bufio.NewScanner(f)
		btree   = &BtreeGN[K, T]{}
	)

	for scanner.Scan() {

		line := scanner.Bytes()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("scanner.Bytes: %v", err)
		}

		line = bytes.TrimSuffix(line, []byte("\n"))

		var entry exportEntry[K, T]
		err = json.Unmarshal(line, &entry)
		if err != nil {
			return nil, fmt.Errorf("json.Unmarshal: %v", err)
		}

		n := btree.Insert(entry.Key, entry.Payload)
		if n == nil {
			return nil, fmt.Errorf("btree.Insert: %v", err)
		}
	}

	return btree, nil
}
