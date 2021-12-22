package main

import (
	"fmt"
	"time"
)

func running() {
	var times int
	for {
		times++
		fmt.Println("tick", times)
		time.Sleep(time.Second)
	}
}

func main() {
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
