package hashing

import (
	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
)

type Hashtable[K comparable, V any] struct {
	elems []*bucket[K, V]
	hash  utils.Hasher[K]
	size  int
}

type bucket[K any, V any] struct {
	head *bucketNode[K, V]
}

type bucketNode[K any, V any] struct {
	key   K
	value V
	next  *bucketNode[K, V]
}

// New initializes the hashtable passed.
func (h *Hashtable[K, V]) New(hash func(K, int) int, size int) *Hashtable[K, V] {
	h = &Hashtable[K, V]{elems: make([]*bucket[K, V], size), size: size, hash: hash}
	for i := range h.elems {
		h.elems[i] = &bucket[K, V]{}
	}
	return h
}

// Insert into the hashtable given key and value.
// returns uninitialized error if hashtable is not initialized.
func (h Hashtable[K, V]) Insert(key K, val V) error {
	if len(h.elems) == 0 || h.elems == nil {
		return errors.UninitializedError
	}
	hash := h.hash(key, h.size)
	index := hash % h.size
	h.elems[index].insert(key, val)
	return nil
}

func (b *bucket[K, V]) insert(key K, val V) {
	newNode := &bucketNode[K, V]{key: key, value: val}
	newNode.next = b.head
	b.head = newNode
}

// Search in hashtable given the key.
// returns NotFound error if key is not found and uninitialized error if hashtable is not initialized.
func (h Hashtable[K, V]) Search(key K) (V, error) {
	var ret V
	if len(h.elems) == 0 || h.elems == nil {
		return ret, errors.UninitializedError
	}
	i := h.hash(key, h.size)
	bucket := h.elems[i]
	node := bucket.head
	for node != nil {
		if node.key == key {
			return node.value, nil
		}
		node = node.next
	}
	return ret, errors.NotFound
}

// Remove from hashtable
func (h *Hashtable[K, V]) Remove(key K) error {
	if len(h.elems) == 0 || h.elems == nil {
		return errors.UninitializedError
	}
	i := h.hash(key, h.size)
	bucket := h.elems[i]
	if bucket.head != nil {
		if bucket.head.key == key {
			bucket.head = bucket.head.next
			return nil
		}
	} else {
		return errors.NotFound
	}

	previousNode := bucket.head
	for previousNode.next != nil {
		if previousNode.next.key == key {
			previousNode.next = previousNode.next.next
		}
		previousNode = previousNode.next
	}
	return nil
}
