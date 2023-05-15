package main

import "fmt"

func main() {
	fmt.Println("tom\tjack")
	fmt.Println("Hello\nWorld")
	fmt.Println("/etc/yum.repos.d")
	fmt.Println("c:\\Users\\Administrator\\Desktop")
	fmt.Println("tom say \"Hello Golang\"")
	// \r 回车，从当前行的最前面开始输出，覆盖掉以前的内容，结果为Golangorld,ilove
	fmt.Println("HelloWorld,ilove\rGolang")
	fmt.Println("HelloWorld,ilove\rGolang")
}
