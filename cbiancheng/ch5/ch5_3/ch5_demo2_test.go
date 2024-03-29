package ch5_3

import (
	"fmt"
	"testing"
)

type Data struct {
	complax  []int
	instance InnerData
	ptr      *InnerData
}

type InnerData struct {
	a int
}

// 值传递测试函数
func passByValue(inFunc Data) Data {
	// 输出参数的成员情况
	fmt.Printf("inFunc value: %+v\n", inFunc)
	// 打印inFunc的指针
	fmt.Printf("inFunc ptr: %p\n", &inFunc)
	inFunc.instance.a = 888
	inFunc.ptr.a = 222
	return inFunc
}

func TestPassByValue(t *testing.T) {
	in := Data{
		complax:  []int{1, 2, 3},
		instance: InnerData{5},
		ptr:      &InnerData{1},
	}

	fmt.Printf("in value: %+v\n", in)
	fmt.Printf("in ptr: %p\n", &in)
	fmt.Printf("in ptr value: %v\n", *in.ptr)
	fmt.Println("************")

	out := passByValue(in)

	fmt.Printf("out value: %+v\n", out)
	fmt.Printf("out ptr: %p\n", &out)
	fmt.Printf("out ptr value: %v\n", *out.ptr)
	fmt.Println("************")

	fmt.Printf("in again value: %+v\n", in)
	fmt.Printf("in again ptr: %p\n", &in)
	fmt.Printf("in again ptr value: %v\n", *in.ptr)
}
