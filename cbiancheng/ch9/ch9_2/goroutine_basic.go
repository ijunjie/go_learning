package gorouting_basic__test

import (
	"fmt"
	"testing"
	"time"
)

func running() {
	var times int
	for {
		times++
		fmt.Println("tike", times)
		time.Sleep(time.Second)
	}
}

func TestGoroutine(t *testing.T) {
	go running()
	/*go func() {
		var times int
		for {
			times++
			fmt.Println("tike", times)
			time.Sleep(time.Second)
		}
	}()*/
	var input string
	fmt.Scanln(&input)
}
