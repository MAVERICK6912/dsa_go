package stack

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/stretchr/testify/assert"
)

func TestStackInt(t *testing.T) {
	vals := []int{24, 69, 9, 12, 2, 56}
	var s Stack[int]
	s.New(func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})

	s.Add(vals...)

	s.Push(35)
	val, err := s.Peek()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 35, val)
	assert.Equal(t, len(vals)+1, s.Size())

	s.Pop()
	val, err = s.Peek()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 56, val)
	assert.Equal(t, len(vals), s.Size())
}

func TestStackString(t *testing.T) {
	vals := []string{"shazam", "green arrow", "wonder woman"}
	var s Stack[string]
	s.New(func(a, b string) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})

	s.Add(vals...)

	s.Push("aquaman")
	val, err := s.Peek()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "aquaman", val)
	assert.Equal(t, len(vals)+1, s.Size())

	val, err = s.Pop()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "aquaman", val)
	assert.Equal(t, len(vals), s.Size())

	val, err = s.Pop()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "wonder woman", val)
	assert.Equal(t, len(vals)-1, s.Size())

	val, err = s.Pop()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "green arrow", val)
	assert.Equal(t, len(vals)-2, s.Size())

	val, err = s.Pop()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "shazam", val)
	assert.Equal(t, 0, s.Size())

	// Popping on empty stack
	val, err = s.Pop()
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.UninitializedError)
	}
	assert.Equal(t, 0, s.Size())
}
