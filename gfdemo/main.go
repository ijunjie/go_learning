package main

import (
	_ "gfdemo/boot"
	_ "gfdemo/router"
	"github.com/gogf/gf/frame/g"
)

func main() {
	g.Server().Run()
}
