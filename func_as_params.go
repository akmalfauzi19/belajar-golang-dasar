package main

type Filter func(string) string

func sayHelloWithFillter(name string, filter Filter) {
	nameFiltered := filter(name)
	println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFillter("anjing", spamFilter)

	filter := spamFilter
	sayHelloWithFillter("budi", filter)
}
