package mutext_test

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

var counter int = 0

func Count(lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	counter++
	fmt.Println(counter)
}

func TestMutex(t *testing.T) {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go Count(lock)
	}
	for {
		lock.Lock()
		c := counter
		lock.Unlock()
		runtime.Gosched()
		if c >= 10 {
			break
		}
	}
}
