package ch9_8k

import "testing"

// 运行时发现所有的 goroutine（包括main）都处于等待 goroutine。也就是说所有 goroutine 中的 channel 并没有形成发送和接收对应的代码。
func TestChann(t *testing.T) {
	ch1 := make(chan int)
	ch1 <- 0

	ch2 := make(chan interface{})
	ch2 <- "hello"

	type Equip struct {
		Name string
	}

	ch3 := make(chan *Equip)
	ch3 <- &Equip{
		"hello",
	}
}
