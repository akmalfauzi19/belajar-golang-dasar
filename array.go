package main

import "fmt"

func main() {
	var name [3]string

	name[0] = "Hello"
	name[1] = "World"
	name[2] = "!"

	fmt.Println(name[0], name[1], name[2])

	// Array with values

	var number = [3]int{
		1, 2, 3,
	}

	fmt.Println(number[0], number[1], number[2])

	// array with ellipsis
	// penggunaan [...] untuk menentukan panjang array
	var number2 = [...]int{
		1, 2, 3,
	}

	fmt.Println(number2[0], number2[1], number2[2])

	// array with index

	var number3 = [3]int{
		0: 1,
		1: 2,
		2: 3,
	}

	fmt.Println(number3[0], number3[1], number3[2])
}
