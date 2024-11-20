package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name

	// solusinya
	// menambahkan * pada parameter
}

func main() {
	akmal := Man{"Akmal"}
	akmal.Married()

	fmt.Println(akmal.Name)

	// note:
	// jika Married() menggunakan pointer receiver, maka outputnya adalah "Mr. Akmal"
	// tapi karena Married() menggunakan value receiver, maka outputnya adalah "Akmal"
	// akmal.Married() tidak akan merubah nilai dari akmal.Name
	// karena method Married() menerima value receiver, bukan pointer receiver
	// sehingga, ketika Married() dipanggil, maka akan dibuat copy dari akmal
	// dan copy tersebut yang akan diubah

}
