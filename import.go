package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("Akmal")
	fmt.Println(result)

	fmt.Println(helper.Application)
	// fmt.Println(helper.version) // error
	// fmt.Println(helper.sayGoodBye("Akmal")) // error

	// jika ingin mengakses fungsi private, maka harus membuat fungsi public yang mengakses fungsi private tersebut
}
