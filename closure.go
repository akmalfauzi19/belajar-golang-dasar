package main

func main() {
	counter := 0

	increment := func() {
		println("Incrementing counter")
		counter++
	}

	increment()
	increment()
	increment()

	println(counter)

	// Output:
	// Incrementing counter
	// Incrementing counter
	// Incrementing counter
	// 3

}
