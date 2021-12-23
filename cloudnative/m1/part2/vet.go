package main

import (
	"log"
	"time"
)

func main() {
	words := []string{"foo", "bar", "baz"}
	for _, word := range words {
		// 否则会全部打印 baz
		x := word
		go func() {
			log.Print(x)
		}()
	}
	time.Sleep(2 * time.Second)
}
