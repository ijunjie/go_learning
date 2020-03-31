package main

import "fmt"

func main() {
	// 示例 1
	ch1 := make(chan int, 1)
	ch1 <- 1
	fmt.Println(<-ch1)
	ch1 <- 2
	fmt.Println(<-ch1)
	// 示例 2
	ch2 := make(chan int, 1)
	ch2 <- 1
	elem, ok := <-ch2
	_,_ = elem, ok

	// 示例3。
	var ch3 chan int
	//ch3 <- 1 // 通道的值为nil，因此这里会造成永久的阻塞！
	//<-ch3 // 通道的值为nil，因此这里会造成永久的阻塞！
	_ = ch3
}
