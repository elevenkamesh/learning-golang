package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("welcome to math is golangs")
	var numOne int = 2
	var numTwo float64 = 4

	fmt.Print(numOne + int(numTwo))
	// rand.Seed(time.Now().UnixNano())
	// fmt.Print(rand.Intn(5))
	myrandNum, _ := rand.Int(rand.Reader, big.NewInt(5))
	print(myrandNum)
}
