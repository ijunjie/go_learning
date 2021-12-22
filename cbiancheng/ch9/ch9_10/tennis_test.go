package ch9_10

import (
	"log"
	"math/rand"
	"sync"
	"testing"
)

var wg sync.WaitGroup

func TestTennis(t *testing.T) {
	court := make(chan int)
	wg.Add(2)
	go player("Nadal", court)
	go player("Djokovic", court)

	court <- 1
	wg.Wait()
}

func player(name string, court chan int) {
	defer wg.Done()

	for {
		ball, ok := <-court
		if !ok {
			log.Printf("Player %s Won\n", name)
			return
		}

		n := rand.Intn(100)
		if n%13 == 0 {
			log.Printf("Player %s missed\n", name)
			close(court)
			return
		}
		log.Printf("Player %s Hit %d\n", name, ball)
		ball++
		court <- ball
	}
}
