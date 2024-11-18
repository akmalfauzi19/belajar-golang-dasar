package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "wick",
		"address": "martapura",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["title"] = "belajar golang"
	book["author"] = "wick"
	book["ups"] = "salah"

	fmt.Println(book)

	delete(book, "ups")
	fmt.Println(book)

}
