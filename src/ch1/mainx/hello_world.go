package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
	fmt.Println("Hello World!")
	os.Exit(0)
}

// go run
// go build 会在当前路径下生成可执行文件

// 程序入口只有两点要求： package main 和 func main
// 目录和源码文件名不一定是main

// main的返回值通过 os.Exit
// 参数通过 os.Args 传入 Args 首个元素是程序本身
