package concurrency

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	fmt.Println("service...1")
	time.Sleep(time.Millisecond * 50)
	fmt.Println("service...2")
	return "Done"
}

func otherTask() {
	fmt.Println("other task...1")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("other task...2")
}

func AsyncService() chan string {
	//retCh := make(chan string) //只有 chan 被取出，才输出 after channel
	retCh := make(chan string, 1) // 写入 chan 后立即输出 after channel
	go func() {
		ret := service()
		fmt.Println("before channel")
		retCh <- ret
		fmt.Println("after channel")
	}()
	return retCh
}

// func TestService(t *testing.T) {
// 	fmt.Println(service())
// 	otherTask()
// }

func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	otherTask()
	fmt.Println(<-retCh)
	time.Sleep(time.Second * 1)
}
