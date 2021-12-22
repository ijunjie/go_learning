package ch9_9k

import (
	"fmt"
	"testing"
)

func printer(c chan int) {
	// 开始无限循环等待数据
	for {
		// 从channel中获取一个数据
		data := <-c
		// 将 -1 视为数据结束
		if data == -1 {
			break
		}
		// 打印数据
		fmt.Println(data)
	}

	// 通知 main已经结束循环
	c <- 204
}

func printer2(c chan int) {

	for data := range c {
		if data == -1 {
			break
		}
		fmt.Println(data)
	}

	// 通知main已经结束循环（我搞定了！）
	c <- 204

}

func TestPrinter1(t *testing.T) {
	// 创建一个 channel
	c := make(chan int)
	// 并发执行 printer，传入 channel
	go printer(c)

	for i := 1; i <= 10; i++ {
		// 将数据通过 channel 投送给 printer
		c <- i
	}

	// 通知并发的 printer 结束循环（没数据啦！）
	c <- -1

	// 阻塞等待 printer 结束（搞定喊我！）
	r := <-c
	fmt.Printf("result %d\n", r)
}

func TestPrinter2(t *testing.T) {

	// 创建一个channel
	c := make(chan int)

	// 并发执行printer，传入channel
	go printer2(c)

	for i := 1; i <= 10; i++ {

		// 将数据通过channel投送给printer
		c <- i
	}

	// 通知并发的printer结束循环（没数据啦！）
	c <- -1

	// 等待printer结束（搞定喊我！）
	r := <-c
	fmt.Printf("result %d\n", r)
}
