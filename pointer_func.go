package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	address := Address{}
	ChangeCountryToIndonesia(&address)
	fmt.Println(address.Country) // Indonesia

	/*
		penjelasan.
		1. address adalah struct dari Address
		2. ChangeCountryToIndonesia adalah sebuah fungsi yang menerima pointer dari Address
		3. &address adalah pointer dari address
		4. jadi, ketika kita mempassing &address ke ChangeCountryToIndonesia, maka itu akan mengubah nilai dari address.Country

	*/
}
