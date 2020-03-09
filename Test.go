package main

import (
	"fmt"
	"unsafe"
)

var (
	value1 = 1
	value2 = 2
)

func main() {
	var a int = 9223372036854775807

	fmt.Println(unsafe.Sizeof(a))

	fmt.Println(len("a"))

	fmt.Println(a, value1, value2)

	fmt.Println("Hello World!" + " GO")

	const (
		i = 1 << iota
		j = 3 << iota
		k
		l
	)

	fmt.Println(l)

	//ConsolePrintln("GO!")
}
