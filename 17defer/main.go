package main

import "fmt"

func main() {

	fmt.Println("Hello World!")
	defer fmt.Println("world")
	defer fmt.Println("ones")

	defer fmt.Println("two")
	defer fmt.Println("three")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
