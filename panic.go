package main

func endApp() {
	// message := recover()
	// if message != nil {
	// 	println("error message: ", message)
	// }

	println("Aplikasi selesai")
	message := recover()
	println("terjadi panic: ", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("APLIKASI ERROR")
	}
	println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	println("Aplikasi berjalan")
}
