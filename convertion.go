package main

import "fmt"

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16) // hasil negatif karena melebihi kapasitas int16

	// convertion string
	var name = "john wick"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(eString)
}
