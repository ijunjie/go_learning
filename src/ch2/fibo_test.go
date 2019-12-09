package ch2_test

import "testing"

// 使用尾递归实现斐波那契函数

func inner(a, b, i, n int) int {
	if i == n {
		return b
	} else {
		return inner(b, a+b, i+1, n)
	}
}

func fibo(n int) int {
	return inner(0, 1, 0, n)
}

func TestFibo(t *testing.T) {
	t.Log(fibo(0))
	t.Log(fibo(1))
	t.Log(fibo(2))
	t.Log(fibo(3))
	t.Log(fibo(4))
	t.Log(fibo(5))
	t.Log(fibo(6))
}
