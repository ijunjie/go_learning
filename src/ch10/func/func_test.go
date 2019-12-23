package func_test

import "testing"

func TestFunc(t *testing.T) {
	a := sum(1, 2, 3)
	t.Log(a)
	b := sum(1, 2, 3, 4)
	t.Log(b)
}

func sum(ops ...int) int {
	ret := 0
	for _, e := range ops {
		ret += e
	}
	return ret
}
