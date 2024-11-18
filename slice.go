package main

import "fmt"

func main() {
	name := []string{
		"jhon",
		"wick",
		"ethan",
		"hunt",
		"jason",
		"bourne",
	}

	slice1 := name[0:4]
	fmt.Println(slice1)

	slice2 := name[2:5]
	fmt.Println(slice2)

	// tanpa short hand
	var slice3 []string = name[3:]
	fmt.Println(slice3)

	// tanpa short hand
	var slice4 []string = name[:]
	fmt.Println(slice4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:] // ["sabtu", "minggu"]
	fmt.Println(daySlice1)

	daySlice1[0] = "sabtu baru"
	daySlice1[1] = "minggu baru"
	fmt.Println(daySlice1) // ["sabtu baru", "minggu baru"]

	fmt.Println(days) // ["senin", "selasa", "rabu", "kamis", "jumat", "sabtu baru", "minggu baru"]

	// append(daySlice1, "Libur baru") //tidak bisa menyimpan karena kapasitas days sudah penuh

	daySlice2 := append(daySlice1, "Libur baru")
	fmt.Println(daySlice1)
	fmt.Println(daySlice2) // ["sabtu baru", "minggu baru", "Libur baru"]
	fmt.Println(days)      // ["senin", "selasa", "rabu", "kamis", "jumat", "sabtu baru", "minggu baru"]

	// membuat slice baru
	newSlice := make([]string, 2, 5)
	newSlice[0] = "wick"
	newSlice[1] = "hunt"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// appending to newSlice
	newSlice2 := append(newSlice, "bourne", "ethan")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	// copy slice
	newSlice2[0] = "wick baru"
	fmt.Println(newSlice)
	fmt.Println(newSlice2)

	// copy slice
	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)
	fmt.Println(toSlice)
	fmt.Println(fromSlice)

	//  slice vs array

	// array
	array1 := [...]int{1, 2, 3}

	// slice
	slice5 := []int{1, 2, 3}

}
