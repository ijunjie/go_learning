package ch5

import (
	"math"
	"testing"
)

func TestFunc(t *testing.T) {
	t.Log(hypot(3, 4))
	t.Logf("%T\n", add)
	t.Logf("%T\n", sub)
	t.Logf("%T\n", first)
	t.Logf("%T\n", zero)
	a := func() int {
		return 1
	}()
	t.Log(a)
}

func hypot(x, y float64) float64 {
	return math.Sqrt(x*x + y*y)
}

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) (z int) {
	z = x - y
	return
}

func first(x int, _ int) int {
	return x
}

func zero(int, int) int {
	return 0
}









