package main

import "fmt"
func main()  {
	fmt.Println("hello,world")
	println(getSumAndSub(3, 10))
	println(getSumAndSub(5, 10))
}


func getSumAndSub(n1 int, n2 int)(int,int) {
	sum := n1 + n2
	sub := n1 - n2
	return sum, sub
}