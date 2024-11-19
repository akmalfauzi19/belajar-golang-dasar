package main

type HasName interface {
	GetName() string
}

func SayHello(hasName HasName) string {
	return "Hello, " + hasName.GetName()
}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{Name: "John"}
	println(SayHello(person))

	animal := Animal{Name: "Dog"}
	println(SayHello(animal))
}
