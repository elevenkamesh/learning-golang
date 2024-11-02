package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	arr := []string{}
	arr = append(arr, "Hello", "World", "!")
	fmt.Println("array values", arr)
	arr = append(arr, "x", "Wa", "xxq")

	fmt.Println("array 2 values", arr)

	fmt.Println("array slices in slices ", arr[2:5])

	slicex := make([]int, 4)
	fmt.Print(slicex)
	slicex[0] = 4423
	slicex[1] = 223
	slicex[2] = 44423
	slicex[3] = 1
	slicex = append(slicex, 11, 55)
	fmt.Println("after make append", slicex)
	// remove value
	index := 2
	cx := append(slicex[:index], slicex[index+1:]...)
	fmt.Println("after removed index ", cx)

}
