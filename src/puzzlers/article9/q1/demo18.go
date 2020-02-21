package main

func main() {
	// 示例 1
	//var badMap1 = map[[]int]int{} // key 不能是数组
	//_ = badMap1

	// 示例 2
	//var badMap2 = map[interface{}]int {
	//	"1": 1,
	//	[]int{2}: 2, // runtime error: hash of unhashable type []int
	//	3: 3,
	//}
	//_ = badMap2

	// 示例3。
	//var badMap3 map[[1][]string]int // 这里会引发编译错误。
	//_ = badMap3

	// 示例 4
	//type BadKey1 struct {
	//	slice []string
	//}
	//var badMap4 map[BadKey1]int
	//_ = badMap4

	// 示例 5
	//var badMap5 map[[1][2][3][]string]int
	//_ = badMap5

	// 示例6。
	//type BadKey2Field1 struct {
	//	slice []string
	//}
	//type BadKey2 struct {
	//	field BadKey2Field1
	//}
	//var badMap6 map[BadKey2]int // 这里会引发编译错误。
	//_ = badMap6
}
