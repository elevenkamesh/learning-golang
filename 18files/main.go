package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("welcome to file hadling")

	content := "hi this is kamesh from india "
	file, err := os.Create("./kamesh.txt")

	if err != nil {
		panic(err)
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("length is ", length)
	defer file.Close()
	readFile("./kamesh.txt")
}

func readFile(fileName string) {
	datatype, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	fmt.Println("this the file /n", datatype)
	fmt.Println("this the file /n", string(datatype))

}
