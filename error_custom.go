package main

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	if id == "" {
		return &validationError{"ID tidak boleh kosong"}
	}
	if data != "akmal" {
		return &notFoundError{"Data tidak ditemukan"}
	}
	// Simulasi penyimpanan data
	return nil
}

func main() {
	err := SaveData("1", "ak2mal")

	if err != nil {
		switch err.(type) {
		case *validationError:
			println("Error:", err)
		case *notFoundError:
			println("Error:", err)
		default:
			println("Error:", err)
		}
	} else {
		println("Data berhasil disimpan")
	}
}
