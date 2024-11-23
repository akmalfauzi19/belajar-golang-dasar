package helper // sesuai dengan nama folder

var Application = "Golang"
var version = 1.0

// fungsi private
func sayGoodBye(Name string) string {
	return "Good Bye " + Name
}

// fungsi public
func SayHello(Name string) string {
	return "Hello " + Name
}

//  penjelasan
/**
- fungsi sayGoodBye merupakan fungsi private, karena diawali dengan huruf kecil
- fungsi SayHello merupakan fungsi public, karena diawali dengan huruf besar
**/
