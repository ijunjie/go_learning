package pipefilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStraightPipeline_Process(t *testing.T) {
	splitter := NewSplitFilter(",")
	converter := NewToIntFilter()
	sum := NewSumFilter()
	sp := &StraightPipeline{"p1", &[]Filter{splitter, converter, sum}}
	resp, err := sp.Process("1,2,3")

	assert.Nil(t, err)
	assert.Equal(t, 6, resp)
}
