package main

import "fmt"

func main() {
	var address string = "广西南宁"
	fmt.Println(address)

	str1 := `
package main
import (
	"fmt"
	"unsafe"
)

func main() {
	var b = false
	fmt.Println(b)
	fmt.Println(unsafe.Sizeof(b))
	fmt.Println("b 的占用空间 =", unsafe.Sizeof(b))
}
`
	fmt.Println(str1)

	var str2 = "wei" + "min"
	fmt.Println(str2)

}
