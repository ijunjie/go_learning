package string_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"unsafe"
)

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	assert.Equal(t, "", s)
	s = "hello"
	t.Log(len(s))
	//s = "\xE4\xBA\xBB\xFF"
	s = "\xE4\xB8\xA5"
	t.Log(s)
	t.Log(len(s))
	s = "中国"
	t.Log(len(s)) //是 byte数

	t.Log("-----------")

	c := []rune(s)
	t.Log(len(c))
	t.Log("rune size:", unsafe.Sizeof(c[0]))

	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		t.Logf("%[1]c %[1]x", c)
	}
}
