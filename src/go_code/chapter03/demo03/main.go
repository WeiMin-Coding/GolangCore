/*
变量的使用方式
*/
package main

import "fmt"

func main() {
	// 第一种：声明赋值
	var i int
	i = 1
	fmt.Println(i)

	// 第二种；类型推倒
	var num = 10
	var txt = "1234"
	var num2 = 11.1
	fmt.Println(num2, txt, num)

	// 第三种：直接赋值
	name := "tome"
	//name1 := "weimin"
	fmt.Println(name)
}
