package main

import "log"

func main() {
	list := []*Profile{
		&Profile{Name: "张三", Age: 30, Married: true},
		&Profile{Name: "李四", Age: 21},
		&Profile{Name: "王麻子", Age: 21},
	}
	log.Printf("size=%v", len(list))
}

type Profile struct {
	Name    string
	Age     int
	Married bool
}
