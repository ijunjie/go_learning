package ch9_13

import (
	"log"
	"testing"
	"time"
)

// time 包提供的函数 After()，
//从字面意思看就是多少时间之后，
//其参数是 time 包的一个常量，time.Second 表示 1 秒。
//time.After 返回一个通道，这个通道在指定时间后，通过通道返回当前时间。

func TestSelect(t *testing.T) {
	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				log.Printf("num=%d\n", num)
			case x := <-time.After(3 * time.Second):
				log.Printf("timeout, x=%v\n", x)
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	x := <-quit
	if x {
		log.Printf("x=%v", x)
	}
	log.Println("over")
}
