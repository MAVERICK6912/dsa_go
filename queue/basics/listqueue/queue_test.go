package queue

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/stretchr/testify/assert"
)

func TestQueueEnqueueInt(t *testing.T) {
	var q Queue[int]
	q.New(func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})
	q.Enqueue(34)
	assert.Equal(t, 1, q.Size())
	val, err := q.Peek()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 34, val)

	q.Enqueue(45)
	assert.Equal(t, 2, q.Size())
	val, err = q.Peek()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, 34, val)
}

func TestQueueDequeueInt(t *testing.T) {
	vals := []int{7, 8, 9}

	var q Queue[int]
	q.New(func(a, b int) int {
		if a > b {
			return 1
		} else if a < b {
			return -1
		}
		return 0
	})
	q.Enqueue(7)
	q.Enqueue(8)
	q.Enqueue(9)

	val, err := q.Dequeue()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, len(vals)-1, q.Size())
	assert.Equal(t, vals[0], val)

	val, err = q.Dequeue()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, len(vals)-2, q.Size())
	assert.Equal(t, vals[1], val)

	val, err = q.Dequeue()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, len(vals)-3, q.Size())
	assert.Equal(t, vals[2], val)
	// Dequeuing when queue is empty
	_, err = q.Dequeue()
	assert.Equal(t, len(vals)-3, q.Size())
	assert.ErrorIs(t, errors.UninitializedError, err)
}

func TestQueueEnqueueString(t *testing.T) {
	var q Queue[string]
	q.New(func(a, b string) int {
		if len(a) > len(b) {
			return 1
		} else if len(a) < len(b) {
			return -1
		}
		return 0
	})
	q.Enqueue("batman")
	assert.Equal(t, 1, q.Size())
	val, err := q.Peek()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "batman", val)

	q.Enqueue("martian-manhunter")
	assert.Equal(t, 2, q.Size())
	val, err = q.Peek()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, "batman", val)
}

func TestQueueDequeueString(t *testing.T) {
	vals := []string{"batman", "flash", "black cannary"}

	var q Queue[string]
	q.New(func(a, b string) int {
		if len(a) > len(b) {
			return 1
		} else if len(a) < len(b) {
			return -1
		}
		return 0
	})
	q.Enqueue("batman")
	q.Enqueue("flash")
	q.Enqueue("black cannary")

	val, err := q.Dequeue()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, len(vals)-1, q.Size())
	assert.Equal(t, vals[0], val)

	val, err = q.Dequeue()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, len(vals)-2, q.Size())
	assert.Equal(t, vals[1], val)

	val, err = q.Dequeue()
	if err != nil {
		t.Fail()
	}
	assert.Equal(t, len(vals)-3, q.Size())
	assert.Equal(t, vals[2], val)
	// Dequeuing when queue is empty
	_, err = q.Dequeue()
	assert.Equal(t, len(vals)-3, q.Size())
	assert.ErrorIs(t, errors.UninitializedError, err)
}
