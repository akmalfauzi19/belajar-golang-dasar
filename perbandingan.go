package main

import "fmt"

func main() {
	var name1 = "Jhon"
	var name2 = "Doe"

	var result bool = name1 == name2

	fmt.Println(result)

	result2 := name1 != name2
	fmt.Println(result2)
}
