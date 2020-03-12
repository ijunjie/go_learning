package type_test

import "testing"

// alias
type MyInt int64

func TestImplicit(t *testing.T) {
	var a int32 = 1
	var b int64
	//b = a 不支持隐式类型转换
	b = int64(a)
	t.Log(a, b)
	var c MyInt
	// 别名也不能赋值
	// c = b
	c = MyInt(b) // 只能做显式类型
	t.Log(c)
}

func TestPointer(t *testing.T) {
	a := 1
	aPtr := &a
	// aPtr = aPtr + 1 不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestString(t *testing.T) {
	var a string
	t.Log(len(a))
	t.Log(a == "") // golang 中 string 初始值为空字符串
}
