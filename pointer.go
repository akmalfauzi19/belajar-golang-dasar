package main

type Address struct {
	City, Province, Country string
}

func main() {

	//  copy value (false)
	// address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	// address2 := address1 // output

	// pointer
	address1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	address2 := &address1 // & = reference

	address2.City = "Bandung"
	// output
	// {Bandung Jawa Barat Indonesia}
	// symbol & using for pointer (reference)
	println(address1)
	println(address2)
}
