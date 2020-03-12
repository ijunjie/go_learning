package pipefilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSumElems(t *testing.T) {
	sf := NewSumFilter()
	ret, err := sf.Process([]int{1, 2, 3})
	assert.Nil(t, err)
	assert.Equal(t, 6, ret)
}

func TestWrongInputForSumFilter(t *testing.T) {
	sf := NewSumFilter()
	_, err := sf.Process([]float32{1.1, 2.2, 3.1})
	assert.NotNil(t, err)

}
