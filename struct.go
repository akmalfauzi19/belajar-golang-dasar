package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {

	// Create a new customer
	customer := Customer{
		Name:    "John",
		Address: "New York",
		Age:     30,
	}

	// Print the customer details
	fmt.Println(customer)

	budi := Customer{
		Name:    "Budi",
		Address: "Jakarta",
		Age:     25,
	}

	fmt.Println(budi)

	// Call the method
	budi.sayHello("John")
}
