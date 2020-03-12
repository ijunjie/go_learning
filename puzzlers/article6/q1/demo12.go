package main

import "fmt"

var container = []string{"zero", "one", "two"}

func main() {
	container := map[int]string{0: "zero", 1: "one", 2: "two"}

	// 赋给两个变量，如果第二个变量为 false, 则 value 为 zero value
	a, ok1 := interface{}(container).([]string)
	if !ok1 {
		fmt.Println(a == nil)
		fmt.Printf("a is %v\n", a)
		fmt.Print("container is not []string\n")
	}
	// 如果第二个变量为 true, 则 value 为转换为目标类型后的值
	b, ok2 := interface{}(container).(map[int]string)
	if !ok2 {
		fmt.Print("container is not map[int]string\n")
	}
	fmt.Printf("b is %v\n", b)
	fmt.Printf("b type: %T\n", b)
	fmt.Printf("The element is %q. (container type: %T)\n",
		b[1], b)

	// 如果只赋值给一个变量，则断言失败时报 panic
	//c := interface{}(container).([]string)
	//fmt.Printf("c is %v\n", c)

	// 方式 2
	elem, err := getElement(container)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
	fmt.Printf("The element is %q. (container type: %T)\n",
		elem, container)
}

func getElement(container interface{}) (elem string, err error) {
	switch t := container.(type) {
	case []string:
		elem = t[1]
	case map[int]string:
		elem = t[1]
	default:
		err = fmt.Errorf("unsupported container type: %T", container)
		return
	}
	return
}
