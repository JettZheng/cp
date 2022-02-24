package datastructure

import "fmt"

//补码 原码取反 再加1
// 0101
// 1010
// +  1
// ----
// 0001

//负数 为加符号位，负数为1
//10101
//0101
//
//00001
func testbit() {
	i := 13
	print(-(i))
	fmt.Printf("%b\n", -(i))
	print(i & -(i))
}
