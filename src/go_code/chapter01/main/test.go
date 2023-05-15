package main // 一个go文件需要在一个包

import "fmt"

func sayOK() {
	// 输出一段话
	fmt.Println("ok")
}

// 写一个函数，实现同时返回和，差
// go函数支持多个返回值
func getSumAndSub(n1 int, n2 int) (int, int) {
	sum := n1 + n2 // go语句后不需要分号
	sub := n1 - n2
	fmt.Println(sum)
	return sum, sub
}

func main() {
	sayOK()
	getSumAndSub(1, 2)
}
