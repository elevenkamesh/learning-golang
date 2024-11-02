package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	numrat, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	fmt.Println("You entered: ", input)

	if err != nil {
		fmt.Print("failed", err)
	} else {
		fmt.Println("rating : ", numrat)

	}

}
