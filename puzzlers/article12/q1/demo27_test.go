package q1

import "errors"

// 声明 type 的函数作为参数
type operate func(x int, y int) int

func calculate1(x int, y int, op operate) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

// 函数作为参数
func calculate2(x int, y int, op func(int, int) int) (int, error) {
	if op == nil {
		return 0, errors.New("invalid operation")
	}
	return op(x, y), nil
}

// 函数作为返回值
func genCalculator1(op func(int, int) int) func(int, int) (int, error) {
	return func(x int, y int) (int, error) {
		if op == nil {
			return 0, errors.New("invalid operation")
		}
		return op(x, y), nil
	}
}
