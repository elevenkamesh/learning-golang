package main

import "fmt"

func main() {

	hashmap := make(map[string]string)
	hashmap["js"] = "javascript"
	hashmap["py"] = "python"

	fmt.Print(hashmap)
	delete(hashmap, "py")
	fmt.Println(hashmap)

}
