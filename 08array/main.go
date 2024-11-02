package main

import "fmt"

func main() {
	fmt.Print("welcone to array ")
	var fruitList [4]string
	fruitList[0] = "apple"
	fruitList[1] = "banana"
	fruitList[2] = "orange"

	fmt.Print(fruitList)

	var veg = [4]string{"carrot", "potato"}

	fmt.Print(veg)
}
