package bella

import "errors"

var (
	// ErrNilHead head of bst in nil
	ErrNilHead = errors.New("head is nil")
	// ErrNotFound key not found
	ErrNotFound = errors.New("key not found")
)
