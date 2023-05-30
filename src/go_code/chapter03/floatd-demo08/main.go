package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var price float32 = 80.12
	fmt.Println(price)

	var num1 float32 = -0.00080
	var num2 float64 = -78095611.09
	fmt.Println(num1, num2)

	var num3 float32 = -123.00000901
	var num4 float64 = -123.00000901
	fmt.Println(num3, num4)

	num5 := 11.11
	fmt.Printf("%T\n", num5)

	num6 := 5.12
	num7 := .123
	fmt.Printf("%T %T\n", num6, num7)
	fmt.Println(num7)

	c1 := '1'
	fmt.Println(c1, unsafe.Sizeof(num3))
}
