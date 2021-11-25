package ch9_4

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"testing"
)

var (
	count int32
	wg    sync.WaitGroup
)

func TestRace(t *testing.T) {
	wg.Add(2)
	go incCount()
	go incCount()

	wg.Wait()
	fmt.Println(count)
}

func TestAtomic(t *testing.T) {
	wg.Add(2)
	//go incCount()
	//go incCount()
	go incCount2()
	go incCount2()
	wg.Wait()
	fmt.Println(count)
}

func incCount() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		value := count
		runtime.Gosched()
		value++
		count = value
	}
}

func incCount2() {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		atomic.AddInt32(&count, 1)
		runtime.Gosched()
	}
}
