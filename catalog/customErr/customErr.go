package customErr

import "errors"

var (
	ErrNotFound = errors.New("entity not found")
)
