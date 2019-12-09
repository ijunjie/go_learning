package operator_test

import "testing"

func TestCompare(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	b := [...]int{1, 2, 3, 4}
	t.Log(a == b)
}
