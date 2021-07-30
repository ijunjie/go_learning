package func_test

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

// 1. 随机数
// 2. 多返回值
// 3. 高阶函数
// 4. 可变参数
// 5. defer

func returnMultiValue() (int, int) {
	rand.Seed(time.Now().Unix()) // 注意，需要加这一句才能生成随机数
	return rand.Intn(10), rand.Intn(20)
}

func slowFun(op int) int {
	time.Sleep(time.Second * 1)
	return op
}

func timeSpent(inner func(int) int) func(int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent: ", time.Since(start).Seconds())
		return ret
	}
}

func TestFn(t *testing.T) {
	a, b := returnMultiValue()
	t.Log(a, b)

	_ = timeSpent(slowFun)(10)

}

func sum(ops ...int) int {
	ret := 0
	for _, e := range ops {
		ret += e
	}
	return ret
}

func TestSum(t *testing.T) {
	a := sum(1, 2, 3)
	t.Log(a)
	b := sum(1, 2, 3, 4)
	t.Log(b)
}

func clear() {
	fmt.Println("clear resources")
}

func TestDefer(t *testing.T) {
	defer clear()
	fmt.Println("start")
	panic("err")
}
