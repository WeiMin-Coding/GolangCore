package main

import "fmt"

func main() {
	var price float32 = 100.4
	fmt.Println("Price =", price)

	var num1 float32 = -0.000089
	var num2 float64 = -0.000089
	fmt.Println("num1 =", num1)
	fmt.Println("num2 =", num2)

	var num3 float32 = -123.0000901
	var num4 float64 = -123.0000901
	fmt.Println("num3 =", num3)
	fmt.Println("num4 =", num4)
}
