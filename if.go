package main

import "fmt"

func main() {
	name := "wick"

	if name == "wick" {
		fmt.Println("Hello wick")
	} else {
		fmt.Println("Hello")
	}

	// short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
