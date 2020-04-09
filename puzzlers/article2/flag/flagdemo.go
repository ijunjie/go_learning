package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()
	argCount := flag.NArg()
	if argCount != 1 {
		fmt.Println("Usage: flag <md5>\nExample: flag c4ca4238a0b923820dcc509a6f75849b")
		return
	}
	arg := flag.Arg(0)
	fmt.Println(arg)
}
