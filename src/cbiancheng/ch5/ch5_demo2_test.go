package ch5

import (
	"fmt"
	"testing"
)

type Data struct {
	complax []int
	instance InnerData
	ptr *InnerData
}

type InnerData struct {
	 a int
}

func passByValue(inFunc *Data) *Data {
	fmt.Printf("inFunc value: %+v\n", *inFunc)
	fmt.Printf("inFunc ptr: %p\n", inFunc)
	inFunc.complax = []int{22}
	return inFunc
}


func TestPassByValue(t *testing.T) {
	in := Data {
		complax: []int{1,2,3},
		instance: InnerData{5},
		ptr: &InnerData{1},
	}

	t.Logf("in value: %+v\n", in)
	t.Logf("in ptr: %p\n", &in)


	out := passByValue(&in)

	t.Logf("out value: %+v\n", *out)
	t.Logf("out ptr: %p\n", out)


	t.Logf("in again value: %+v\n", in)
	t.Logf("in again ptr: %p\n", &in)
}
