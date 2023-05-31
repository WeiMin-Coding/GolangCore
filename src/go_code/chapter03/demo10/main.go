package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println(b)
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b))
}
