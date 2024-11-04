package main

import "fmt"

func main() {
	var name string

	name = "John Doe"
	fmt.Println(name)

	name = "John Wick"
	fmt.Println(name)

	// short hand hanya bisa digunakan deklasari pertama kali
	name2 := "John Wick"
	fmt.Println(name2)

	name2 = "John Doe"
	fmt.Println(name2)

	// deklarasi multiple variable
	var (
		firstName = "John"
		lastName  = "Doe"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}
