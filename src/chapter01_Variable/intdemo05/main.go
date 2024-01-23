package main

import (
	"fmt"
	"unsafe"
)

// go中整数类型的使用
func main() {
	var i int = 1
	fmt.Println(i)

	// 测试int8范围
	var i1 int8 = -128
	fmt.Println("i1 =", i1)

	fmt.Printf("n1 Type %T\n", i1)

	var n2 int64 = 10
	fmt.Printf("n2 Type %T, n2 Size %d", n2, unsafe.Sizeof(n2))
}
