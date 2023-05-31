package main

import "fmt"

func main() {
	var c1 byte = 'a'
	var c2 byte = '0'
	fmt.Println(c1, c2)
	fmt.Printf("%c %c\n", c1, c2)

	var c3 int = '韦'
	fmt.Printf("%c %d\n", c3, c3)

	var c4 int = 22269
	fmt.Printf("%c\n", c4)

	var n1 = 10 + 'a'
	fmt.Println(n1)
}
