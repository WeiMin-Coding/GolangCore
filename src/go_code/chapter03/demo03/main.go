/*
变量的使用方式
*/
package main

import "fmt"

// 全局变量声明方式一
var x1 = 300
var x2 = 200
var name3 = "tom"

// 全局变量声明方式二
var (
	x3    = 300
	x4    = 300
	name4 = "tom"
)

func main() {
	// 第一种
	var n1, n2, n3 int
	fmt.Println(n1, n2, n3)

	// 第二种
	var n11, name, n31 = 100, "weimin", 55
	fmt.Println(n11, name, n31)

	// 第三种
	n100, name2, n300 := 100, "weimin", 1000
	fmt.Println(n100, name2, n300)

	// 输出全局变量
	fmt.Println(x1, x2, name3)
	fmt.Println(x3, x4, name4)
}
