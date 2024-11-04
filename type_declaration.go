package main

import "fmt"

func main() {
	type NoKTP string

	var ktpJhon NoKTP = "12313213"
	fmt.Println(ktpJhon)

	var contohString string = "123123"
	var contohKtp = NoKTP(contohString)

	fmt.Println(contohKtp)
}
