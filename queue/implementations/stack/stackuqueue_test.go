package queue

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

func TestStackUQueuePush(t *testing.T) {
	var s *Stack[int]
	s = s.New(utils.IntComparator)

	err := s.Push(20)
	if err != nil {
		t.Error("error while pushing to stack, error was: ", err.Error())
	}
	actual, err := s.Peek()
	if err != nil {
		t.Error("error while getting from stack, error was: ", err.Error())
	}
	assert.Equal(t, 20, actual)
	assert.Equal(t, 1, s.Size())

	err = s.Push(69)
	if err != nil {
		t.Error("error while pushing to stack, error was: ", err.Error())
	}
	actual, err = s.Peek()
	if err != nil {
		t.Error("error while getting from stack, error was: ", err.Error())
	}
	assert.Equal(t, 69, actual)
	assert.Equal(t, 2, s.Size())

	err = s.Push(9)
	if err != nil {
		t.Error("error while pushing to stack, error was: ", err.Error())
	}
	actual, err = s.Peek()
	if err != nil {
		t.Error("error while getting from stack, error was: ", err.Error())
	}
	assert.Equal(t, 9, actual)
	assert.Equal(t, 3, s.Size())
}

func TestStackUQueuePop(t *testing.T) {
	var s *Stack[int]
	s = s.New(utils.IntComparator)

	err := s.Push(125)
	if err != nil {
		t.Error("error while pushing to stack, error was: ", err.Error())
	}
	err = s.Push(2169)
	if err != nil {
		t.Error("error while pushing to stack, error was: ", err.Error())
	}
	err = s.Push(3)
	if err != nil {
		t.Error("error while pushing to stack, error was: ", err.Error())
	}

	actual, err := s.Pop()
	if err != nil {
		t.Error("error while peeking from stack, error was: ", err.Error())
	}
	assert.Equal(t, 3, actual)
	assert.Equal(t, 2, s.Size())

	actual, err = s.Pop()
	if err != nil {
		t.Error("error while getting from stack, error was: ", err.Error())
	}
	assert.Equal(t, 2169, actual)
	assert.Equal(t, 1, s.Size())

	actual, err = s.Pop()
	if err != nil {
		t.Error("error while getting from stack, error was: ", err.Error())
	}
	assert.Equal(t, 125, actual)
	assert.Equal(t, 0, s.Size())

	// popping when stack is empty
	_, err = s.Pop()
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.UninitializedError)
	}
	assert.Equal(t, 0, s.Size())
}
