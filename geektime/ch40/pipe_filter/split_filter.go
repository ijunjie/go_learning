package pipefilter

import (
	"errors"
	"strings"
)

var SplitFilterWrongFormatError = errors.New("input data should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter}
}

func (sf *SplitFilter) Process(request Request) (Response, error) {
	str, ok := request.(string)
	if !ok {
		return nil, SplitFilterWrongFormatError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}
