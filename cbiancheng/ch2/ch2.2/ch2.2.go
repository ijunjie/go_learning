package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17 // 如果不指定 float32，Go语言编译器会将 damageRate 类型推导为 float64
	var damage = float32(attack-defence) * damageRate
	fmt.Println(damage)

	dial, err := net.Dial("tcp", "127.0.0.1:8080")
	if err == nil {
		log.Fatal(err)
	}
	if dial == nil {
		fmt.Println("dial is nil")
		os.Exit(0)
	}
	addr := dial.RemoteAddr()
	fmt.Println(addr)
}
