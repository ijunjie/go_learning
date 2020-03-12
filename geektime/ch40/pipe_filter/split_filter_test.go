package pipefilter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringSpit(t *testing.T) {
	sf := NewSplitFilter(",")
	resp, err := sf.Process("1,2,3")
	assert.Nil(t, err)

	parts, ok := resp.([]string)
	assert.True(t, ok)
	assert.EqualValues(t, []string{"1", "2", "3"}, parts)
}

func TestWrongInput(t *testing.T) {
	sf := NewSplitFilter(",")
	_, err := sf.Process(123)
	assert.NotNil(t, err)
}
