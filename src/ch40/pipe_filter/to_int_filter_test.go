package pipefilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertToInt(t *testing.T) {
	tif := NewToIntFilter()
	resp, err := tif.Process([]string{"1", "2", "3"})
	assert.Nil(t, err)
	assert.EqualValues(t, []int{1, 2, 3}, resp)
}

func TestWrongInputForTIF(t *testing.T) {
	tif := NewToIntFilter()
	_, err := tif.Process([]string{"1", "2.2", "3"})
	assert.NotNil(t, err)
}
