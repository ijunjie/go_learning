package main

import (
	"fmt"
	"testing"
)

// 声明一个函数
type Printer func(contents string) (n int, err error)

// 定义实现
func printToStd(contents string) (n int, err error) {
	return fmt.Println(contents)
}

func TestFunc(t *testing.T) {
	var p Printer
	p = printToStd
	_, _ = p("something")
}