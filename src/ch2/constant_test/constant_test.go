package const_test

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

const (
	Readable = 1 << iota
	Writable
	Executable
)

func TestConst(t *testing.T) {
	t.Log(Monday)
	// t.Log(Tuesday)
	// t.Log(Wednesday)
	// t.Log(Thursday)
	// t.Log(Friday)
	// t.Log(Saturday)
	// t.Log(Sunday)
	a := 7
	b := 1
	t.Log(a&Readable, a&Writable, a&Executable)
	t.Log(b&Readable, b&Writable, b&Executable)
}
