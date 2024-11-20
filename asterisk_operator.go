package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {
	var address1 Address = Address{"Jakarta", "DKI", "Indonesia"}
	var address2 *Address = &address1

	address2.City = "Pekalongan"

	fmt.Println(address1)
	fmt.Println(address2)

	// address2 = &Address{"solo", "Jawa tengah", "Indonesia"}
	// fmt.Println("=====================================")
	// fmt.Println(address1)
	// fmt.Println(address2)

	// output
	// {Pekalongan DKI Indonesia}
	// &{Pekalongan DKI Indonesia}

	// jika ingin mengubah address1 juga maka gunakan *address2
	*address2 = Address{"Bandung", "Jawa Barat", "Indonesia"}
	println("=====================================")
	fmt.Println(address1)
	fmt.Println(address2)
	// output
	// {Bandung Jawa Barat Indonesia}

}
