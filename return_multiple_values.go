package main

import "fmt"

func getFullName() (string, string) {
	return "John", "Doe"
}

func main() {
	firstName, lastName := getFullName()
	fmt.Println("Hello", firstName, lastName)

	// if use _ to ignore the return value
	_, lastName2 := getFullName()
	fmt.Println("Hello", lastName2)
}
