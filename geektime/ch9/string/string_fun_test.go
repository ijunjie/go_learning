package string_test

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

func TestStringFunc(t *testing.T) {
	s := "A,B,C"
	parts := strings.Split(s, ",")
	assert.EqualValues(t, []string{"A", "B", "C"}, parts)

	joined := strings.Join(parts, "-")
	assert.EqualValues(t, "A-B-C", joined)
}

func TestStrConf(t *testing.T) {
	s := strconv.Itoa(10)
	assert.Equal(t, "10", s)

	i, err := strconv.Atoi("20")
	assert.Nil(t, err)
	assert.Equal(t, 20, i)

	i2, err2 := strconv.Atoi("2xx")
	assert.NotNil(t, err2)
	assert.Equal(t, 0, i2)
}
