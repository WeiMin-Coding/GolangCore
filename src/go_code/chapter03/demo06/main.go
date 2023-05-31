package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int = 1
	fmt.Println(i)

	var j int8 = -127
	fmt.Println(j)

	var x uint8 = 255
	fmt.Println(x)

	var n1 = 100
	var b uint = 1
	var c byte = 255
	fmt.Println(n1, b, c)
	fmt.Printf("n1 type %T\n", n1)

	var n2 int64 = 10
	fmt.Printf("n2 %d", unsafe.Sizeof(n2))
}
