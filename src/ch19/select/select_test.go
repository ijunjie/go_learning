package select_test

import (
	"fmt"
	"testing"
	"time"
)

func async1() chan string {
	retCh := make(chan string, 1)
	go func() {
		time.Sleep(time.Millisecond * 50)
		retCh <- "async1"
		fmt.Println("finish async1")
	}()
	return retCh
}

func async2() chan string {
	retCh := make(chan string, 1)
	go func() {
		time.Sleep(time.Millisecond * 80)
		retCh <- "async2"
		fmt.Println("finish async2")
	}()
	return retCh
}

// func TestSelect(t *testing.T) {
// 	select {
// 	case ret := <-async1():
// 		t.Logf("result %s", ret)
// 	case ret := <-async2():
// 		t.Logf("result %s", ret)
// 		// default:
// 		// 	t.Error("No one returned")
// 	}
// 	time.Sleep(time.Second * 2)
// }

func TestTimeout(t *testing.T) {
	select {
	case ret := <-async2():
		t.Logf("result %s", ret)
	case <-time.After(time.Millisecond * 100):
		t.Error("timeout")
	}
}
