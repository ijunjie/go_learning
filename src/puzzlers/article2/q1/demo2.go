package main

import (
	"flag"
	"fmt"
)

var name string

func init() {
	flag.StringVar(&name, "name", "everyone", "The greeting object.")
}

func main() {
	flag.Usage = func() {

	}
	flag.Parse()
	fmt.Printf("Hello, %s!\n", name)
}
