package sort

import (
	"testing"

	"github.com/maverick6912/dsa_go/errors"
	"github.com/maverick6912/dsa_go/utils"
	"github.com/stretchr/testify/assert"
)

func TestBubbleSort(t *testing.T) {
	actual, err := BubbleSort(testDataIntAsc.elems, testDataIntAsc.sd, utils.IntComparator)
	if err != nil {
		t.Errorf("error while sorting %v, error was: %v", testDataIntAsc.elems, err)
	}
	assert.Equal(t, testDataIntAsc.expected, actual)

	actual, err = BubbleSort(testDataIntDesc.elems, testDataIntDesc.sd, utils.IntComparator)
	if err != nil {
		t.Errorf("error while sorting %v, error was: %v", testDataIntDesc.elems, err)
	}
	assert.Equal(t, testDataIntDesc.expected, actual)

	actual, err = BubbleSort(testDataIntAscRepeatElem.elems, testDataIntAscRepeatElem.sd, utils.IntComparator)
	if err != nil {
		t.Errorf("error while sorting %v, error was: %v", testDataIntAscRepeatElem.elems, err)
	}
	assert.Equal(t, testDataIntAscRepeatElem.expected, actual)

	_, err = BubbleSort([]string{}, Ascending, utils.StringComparator)
	if assert.Error(t, err) {
		assert.ErrorIs(t, err, errors.NoElements)
	}
}
