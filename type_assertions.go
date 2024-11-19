package main

import "fmt"

func random() any {
	return 10
}

func main() {
	var result = random()
	// var resultString string = result.(string) // panic because result is not string
	// fmt.Println(resultString)

	// using type assertions with switch
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")

	}
}
