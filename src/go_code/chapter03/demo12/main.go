package main

import "fmt"

func main() {
	var i1 int32 = 100
	var f1 float32 = float32(i1)
	var i2 int8 = int8(i1)
	var i3 int64 = int64(i1)
	fmt.Println(i1, f1, i2, i3)

	var num1 int64 = 99999999
	var num2 int8 = int8(num1)
	fmt.Println(num1, num2)
}
