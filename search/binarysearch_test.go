package search

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/sort"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

var (
	testDataIntDesc = struct {
		elems       []int
		searchElems []int
		expected    []int
	}{
		elems:       []int{2160, 1440, 1080, 720, 480, 360, 240},
		searchElems: []int{720, 360, 1440, 0},
		expected:    []int{3, 5, 1, -1},
	}
)

func TestBinarySearch(t *testing.T) {
	for i, v := range testDataInt.searchElems {
		actual, err := BinarySearch(testDataInt.elems, v, sort.Ascending, utils.IntComparator)
		if i == len(testDataInt.searchElems)-1 {
			if assert.Error(t, err) {
				assert.ErrorIs(t, err, errors.NotFound)
			}
			continue
		}
		if err != nil {
			t.Errorf("error while searching for %d, error was: %v", v, err)
		}
		assert.Equal(t, testDataInt.expected[i], actual)
	}

	for i, v := range testDataIntDesc.searchElems {
		actual, err := BinarySearch(testDataIntDesc.elems, v, sort.Descending, utils.IntComparator)
		if i == len(testDataIntDesc.searchElems)-1 {
			if assert.Error(t, err) {
				assert.ErrorIs(t, err, errors.NotFound)
			}
			continue
		}
		if err != nil {
			t.Errorf("error while searching for %d, error was: %v", v, err)
		}
		assert.Equal(t, testDataIntDesc.expected[i], actual)
	}
}
