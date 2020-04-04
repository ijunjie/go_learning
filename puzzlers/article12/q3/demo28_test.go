package q3

import (
	"fmt"
	"testing"
)

// 数组是值类型，传参是拷贝的过程
func TestFunc1(t *testing.T) {
	array1 := [3]string{"a", "b", "c"}
	fmt.Printf("The array: %v\n", array1)
	array2 := modifyArray(array1)
	fmt.Printf("The modified array: %v\n", array2)
	fmt.Printf("The original array: %v\n", array1)
	fmt.Println()
}

func modifyArray(a [3]string) [3]string {
	a[1] = "x"
	return a
}

// 对于引用类型，比如：切片、字典、通道，像上面那样复制它们的值，只会拷贝它们本身而已
// 并不会拷贝它们引用的底层数据。
func TestFunc2(t *testing.T) {
	slice1 := []string{"x", "y", "z"}
	fmt.Printf("The slice: %v\n", slice1)
	slice2 := modifySlice(slice1)
	fmt.Printf("The modified slice: %v\n", slice2)
	fmt.Printf("The original slice: %v\n", slice1)
	fmt.Println()
}

func modifySlice(a []string) []string {
	a[1] = "i"
	return a
}

// 数组中包含切片
func TestFunc3(t *testing.T) {
	complexArray1 := [3][]string{
		[]string{"d", "e", "f"},
		[]string{"g", "h", "i"},
		[]string{"j", "k", "l"},
	}
	fmt.Printf("The complex array: %v\n", complexArray1)
	complexArray2 := modifyComplexArray(complexArray1)
	fmt.Printf("The modified complex array: %v\n", complexArray2)
	fmt.Printf("The original complex array: %v\n", complexArray1)
}

func modifyComplexArray(a [3][]string) [3][]string {
	a[1][1] = "s"
	a[2] = []string{"o", "p", "q"}
	return a
}