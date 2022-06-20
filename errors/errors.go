package errors

import "errors"

var (
	UninitializedError = errors.New("the data-structure was uninitialized")
	IndexOutOfBound    = errors.New("index out of bound")
	NotFound           = errors.New("value not found")
	MissingPriority    = errors.New("priority was nil or empty")
	PeekError          = errors.New("error while peeking")
	DequeueError       = errors.New("error while dequeuing from queue")
	NoElements         = errors.New("no elements to remove")
)
