package main

func main() {
	const name string = "John Doe"
	const age = 30

	/*
		perbedaan const & var, const tidak bisa diubah
		dan const kalau tidak dipakai tidak akan error

	*/

	// error
	// name = "John Wick"
	// age = 35

	// shorthand

	const (
		firstName = "John"
		lastName  = "Doe"
	)

}
