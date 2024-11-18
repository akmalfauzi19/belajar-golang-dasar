package main

func loggin() {
	println("Logging")
}

func runApp() {
	defer loggin()
	println("Running app")
}

func main() {
	runApp()
}
