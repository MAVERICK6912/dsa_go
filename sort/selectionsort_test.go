package sort

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

var (
	testDataIntAsc = struct {
		elems    []int
		expected []int
		sd       SortDirection
	}{
		elems:    []int{2160, 1440, 1080, 720, 480, 360, 240},
		sd:       Ascending,
		expected: []int{240, 360, 480, 720, 1080, 1440, 2160},
	}

	testDataIntAscRepeatElem = struct {
		elems    []int
		expected []int
		sd       SortDirection
	}{
		elems:    []int{2160, 1440, 1080, 720, 480, 360, 240, 1440},
		sd:       Ascending,
		expected: []int{240, 360, 480, 720, 1080, 1440, 1440, 2160},
	}
	testDataIntDesc = struct {
		elems    []int
		expected []int
		sd       SortDirection
	}{
		elems:    []int{240, 360, 480, 720, 1080, 1440, 2160},
		sd:       Descending,
		expected: []int{2160, 1440, 1080, 720, 480, 360, 240},
	}
)

func TestSelectionSort(t *testing.T) {
	actual, err := SelectionSort(testDataIntAsc.elems, testDataIntAsc.sd, utils.IntComparator)
	if err != nil {
		t.Errorf("error while sorting %v, error was: %v", testDataIntAsc.elems, err)
	}
	assert.Equal(t, testDataIntAsc.expected, actual)

	actual, err = SelectionSort(testDataIntDesc.elems, testDataIntDesc.sd, utils.IntComparator)
	if err != nil {
		t.Errorf("error while sorting %v, error was: %v", testDataIntDesc.elems, err)
	}
	assert.Equal(t, testDataIntDesc.expected, actual)

	actual, err = SelectionSort(testDataIntAscRepeatElem.elems, testDataIntAscRepeatElem.sd, utils.IntComparator)
	if err != nil {
		t.Errorf("error while sorting %v, error was: %v", testDataIntAscRepeatElem.elems, err)
	}
	assert.Equal(t, testDataIntAscRepeatElem.expected, actual)

	_, err = SelectionSort([]string{}, Ascending, utils.StringComparator)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NoElements)
	}
}
