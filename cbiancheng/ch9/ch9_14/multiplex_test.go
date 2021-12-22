package ch9_14

import (
	"log"
	"testing"
)

func TestMultiplex(t *testing.T) {
	ch := make(chan int, 1)
	for i := 0; i < 10; i++ {
		select {
		case ch <- 0:
		case ch <- 1:
		}
		i := <-ch
		log.Printf("i=%v\n", i)

	}

}
