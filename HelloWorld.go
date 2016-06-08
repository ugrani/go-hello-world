package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println("hi", addNumber(3, 4))
}

func addNumber(x ,y int) (a int){
	a = x+y
	return
}