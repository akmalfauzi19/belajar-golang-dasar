package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	}
	return nilai / pembagi, nil
}

func main() {
	hasil, err := Pembagian(100, 10)
	if err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println("Hasil:", hasil)
}
