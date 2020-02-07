package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	var a int
	typeOfA := reflect.TypeOf(a)
	t.Log(typeOfA.Name(), typeOfA.Kind())

	type cat struct {
		Name string
		Type int `json:"type" id:"100"`
	}

	aCat := cat{}
	typeOfCat := reflect.TypeOf(aCat)
	t.Log(typeOfCat.Name(), typeOfCat.Kind())
	for i:=0; i<typeOfCat.NumField();i++ {
		fieldType := typeOfCat.Field(i)
		fmt.Printf("name: %v  tag: '%v'\n", fieldType.Name, fieldType.Tag)
	}

	type Enum int
	const (
		Zero Enum = 0
	)
	typeOfZero := reflect.TypeOf(Zero)
	t.Log(typeOfZero.Name(), typeOfZero.Kind())

	catPtr := &aCat
	typeOfCatPtr := reflect.TypeOf(catPtr)
	// 需要注意的是，指针变量的类型名称是空，不是 *cat
	t.Log(typeOfCatPtr.Name(), typeOfCatPtr.Kind())

	elem := typeOfCatPtr.Elem()
	t.Log(elem.Name(), elem.Kind())
}
