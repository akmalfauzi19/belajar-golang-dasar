package main

import "fmt"

func main() {
	name := "John"

	switch name {
	case "John":
		fmt.Println("John")
	case "Doe":
		fmt.Println("Doe")
	default:
		fmt.Println("Default")
	}

	// short statement
	switch name := "Doe"; name {
	case "John":
		fmt.Println("John")
	case "Doe":
		fmt.Println("Doe")
	default:
		fmt.Println("Default")
	}

	// No expression
	switch {
	case name == "John":
		fmt.Println("John")
	case name == "Doe":
		fmt.Println("Doe")
	default:
		fmt.Println("Default")
	}

	// Multiple cases
	switch name {
	case "John", "Doe":
		fmt.Println("John or Doe")
	default:
		fmt.Println("Default")
	}
}
