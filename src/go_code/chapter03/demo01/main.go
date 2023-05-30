package main

import "fmt"

var n1 = 100
var n2 = 200
var name = "jack"

var (
	n3    = 100
	n4    = 900
	name2 = "WeiMin"
)

func main() {
	// 1. 声明变量
	var i int
	// 2. 赋值
	i = 10
	// 3. 使用变量
	fmt.Println(i)

	var num = 10.11
	fmt.Println("num: ", num)

	name := "tom"
	fmt.Println("name: ", name)

	n1, name, n3 := 100, "tom", 888
	fmt.Println(n1, name, n3)
}
