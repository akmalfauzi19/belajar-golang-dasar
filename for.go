package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke ", counter)
		counter++
	}

	fmt.Println("Selesai perulangan")

	for counter1 := 1; counter1 <= 10; counter1++ {
		fmt.Println("Perulangan v2 ke ", counter1)
	}

	fmt.Println("Selesai perulangan v2")

	// Slice
	slice := []string{"satu", "dua", "tiga"}
	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// foreach
	for index, value := range slice {
		fmt.Println("Index", index, "=", value)
	}

	// foreach without index
	for _, value := range slice {
		fmt.Println(value)
	}
}
