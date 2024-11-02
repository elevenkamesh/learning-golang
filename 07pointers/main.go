package main

import "fmt"

func main() {
	fmt.Print("welcome to pointers")
	// var one *int
	// fmt.Println("pointerx=", one)

	number := 12

	ptr := &number

	fmt.Println("value of pointer ", ptr)
	fmt.Println("value of  pointers value  ", *ptr)

}
