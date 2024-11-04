package main

import "fmt"

func main() {

	fmt.Println(len("Hello World"))
	fmt.Println("Hello World"[1]) // result 101 adalah hasil byte
	fmt.Println("Hello " + "World")
	fmt.Println("Hello World"[1:5])

}
