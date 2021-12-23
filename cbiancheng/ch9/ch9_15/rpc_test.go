package ch9_15

import (
	"errors"
	"log"
	"testing"
	"time"
)

func TestRpc(t *testing.T) {
	ch := make(chan string)
	go RPCServer(ch)
	recv, err := Send(ch, "hi")
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("client received %v\n", recv)
}

func Send(ch chan string, req string) (string, error) {
	ch <- req
	select {
	case ack := <-ch:
		return ack, nil
	case <-time.After(time.Second):
		return "", errors.New("Time out")
	}
}

func RPCServer(ch chan string) {
	for {
		data := <-ch
		log.Printf("server received %v\n", data)
		ch <- "roger for " + data
	}
}
