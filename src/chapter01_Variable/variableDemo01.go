package main

import "fmt"

/*
Go变量
1. 局部变量，在方法内。
2. 全局变量，不在方法体内的变量。
3. 声明变量方式，变量都有默认值，如int则是0.
	3.1 var 变量名 = 表达式，如myName1 = "WeiMin"，这种方式会使用类型推倒，自动识别变量类型。
	3.2 var 变量名 类型 = 表达式，如var i int = 10，最常规的方式。
	3.3 变量名 :=表达式，如name := "WeiMin"
	3.4 多变量声明
		声明：var n1, n2, n3 int
		声明并赋值：var n4, n5, n6 = 100, "WeiMin", 1002
*/

// 全局变量
var (
	myName1 = "WeiMin"
	myName2 = 100
	myName3 = 100.1
)

func main() {
	// 局部变量
	// 声明变量方式1
	var myName4 = "WeiMin"

	// 声明变量方式2
	var myName5 string = "WeiMin"

	// 声明变量方式3
	myName6 := 100

	// 声明变量方式4
	var v1, v2, v3 int

	var v4, v5, v6 = 100, "WeiMin", 200.1

	fmt.Println("全局变量：myName1: ", myName1)
	fmt.Println("局部变量：myName4: ", myName4)
	fmt.Println("局部变量：myName5: ", myName5)
	fmt.Println("局部变量：myName5: ", myName6)
	fmt.Println("局部变量：v1, v2, v3: ", v1, v2, v3)
	fmt.Println("局部变量：v4, v5, v6: ", v4, v5, v6)
}
