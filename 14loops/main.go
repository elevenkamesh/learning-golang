package main

import "fmt"

func main() {
	fmt.Println("Starting Loops ")
	days := []string{"mon", "tue", "web", "thur", "fri", "sat", "sun"}
	for i := range days {
		fmt.Println(i, days[i])
	}

	for index, val := range days {
		fmt.Println(index, val)
	}
}
