package mapext

import "testing"

/*
 map 的 value 可以是 func
*/
func TestMapWithFunValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int { return op * op * op }
	t.Log(m[1](2), m[2](2), m[3](2))
}

func TestMapSet(t *testing.T) {
	mySet := map[int]bool{}
	mySet[1] = true
	n := 3
	if v, f := mySet[n]; f {
		t.Logf("%d existing. value is %v", n, v)
	} else {
		t.Logf("%d not existing. value is %v", n, v)
	}
}
