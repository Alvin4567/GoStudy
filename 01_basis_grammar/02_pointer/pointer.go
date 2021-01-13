package main

import "fmt"

//func main() {
//	var i int = 10
//	fmt.Println(&i)
//	var ptr *int = &i
//	fmt.Println(ptr)
//	fmt.Println(&ptr)
//
//	var ptrPtr **int = &ptr
//	fmt.Println(ptrPtr)
//	fmt.Println(&ptrPtr)
//	fmt.Println(*ptrPtr)
//	fmt.Println(*ptr)
//}

func main() {
	var num int = 100
	fmt.Println(num)
	fmt.Println(&num)
	var ptr *int = &num
	fmt.Println(ptr)
	*ptr = 10001
	fmt.Println(num)

	var a, _ = 30, 40
	fmt.Println(a)
}
