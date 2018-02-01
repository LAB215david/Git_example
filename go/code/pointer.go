package main

import "fmt"

func foo(x *int) {
	fmt.Println(x) // function內x的記憶體位置
}

func main() {
	a := 10
	fmt.Println(&a) // main裡面a的記憶體位置
	foo(&a)
}
