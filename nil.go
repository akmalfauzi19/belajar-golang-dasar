package main

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	data := newMap("hello")

	if data == nil {
		println("data is nil")
	} else {
		// call newMap name
		println(data["name"])
	}
}
