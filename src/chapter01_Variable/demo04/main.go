package main

import "fmt"

// go中+号的使用
func main() {
	var i int = 1
	var i1 int = 2
	var i2 int = 3
	var r1 = i + i1 + i2
	fmt.Println("Int + :", r1)

	var str1 string = "hello"
	var str2 string = "world"
	var r2 = str1 + str2
	fmt.Println("String + :", r2)
}
