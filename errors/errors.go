package errors

import "errors"

var (
	UninitializedError = errors.New("the data-structure was uninitialized")
	IndexOutOfBound    = errors.New("index out of bound")
	NotFound           = errors.New("value not found")
)
