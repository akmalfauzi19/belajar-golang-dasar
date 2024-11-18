package main

import "fmt"

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "John"
	middleName = "Wick"
	lastName = "Doe"
	return firstName, middleName, lastName

}

func main() {
	a, b, c := getCompleteName()
	fmt.Println("Hello", a, b, c)
}
