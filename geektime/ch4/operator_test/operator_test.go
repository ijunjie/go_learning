package operator_test

import "testing"

func TestCompare(t *testing.T) {
	// 数组的比较，首先维度不同编译不过
	a := [...]int{1, 2, 3, 5}
	b := [...]int{1, 2, 3, 4}
	t.Log(a == b)
}
