package search

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

var (
	testDataInt = struct {
		elems       []int
		searchElems []int
		expected    []int
	}{
		elems:       []int{240, 360, 480, 720, 1080, 1440, 2160},
		searchElems: []int{720, 360, 1440, 0},
		expected:    []int{3, 1, 5, -1},
	}
	testDataString = struct {
		elems       []string
		searchElems []string
		expected    []int
	}{
		elems:       []string{"240", "360", "480", "720", "1080", "1440", "2160"},
		searchElems: []string{"720", "360", "0"},
		expected:    []int{3, 1, -1},
	}
)

func TestLinearSearch(t *testing.T) {
	for i, v := range testDataInt.searchElems {
		actual, err := LinearSearch(testDataInt.elems, v, utils.IntComparator)
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
	for i, v := range testDataString.searchElems {
		actual, err := LinearSearch(testDataString.elems, v, utils.StringComparator)
		if i == len(testDataString.searchElems)-1 {
			if assert.Error(t, err) {
				assert.ErrorIs(t, err, errors.NotFound)
			}
			continue
		}
		if err != nil {
			t.Errorf("error while searching for %s, error was: %v", v, err)
		}
		assert.Equal(t, testDataString.expected[i], actual)
	}
}
