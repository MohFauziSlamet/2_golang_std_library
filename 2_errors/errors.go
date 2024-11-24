package main

import (
	"errors"
	"fmt"
)

// * memebuat variabel errors
var (
	ValidationErrors = errors.New("Validation Errors")
	NotFoundError    = errors.New("Not Found Error")
)

// * Kode menggunakan error
func GetByID(name string) error {
	if name == "" {
		return ValidationErrors
	}

	if name != "eko" {
		return NotFoundError
	}

	return nil
}

func main() {
	/*
	   Mengecek Jenis Error :
	   Misal kita membuat jenis error sendiri, lalu kita ingin mengecek jenis errornya
	   Kita bisa menggunakan errors.Is() untuk mengecek jenis type error nya
	*/

	err := GetByID("")

	if err != nil {

		if errors.Is(err, ValidationErrors) {
			//* ValidationErrors
			fmt.Println("Validation error")
		} else if errors.Is(err, NotFoundError) {
			//* NotFoundError
			fmt.Println("Not Found Error")
		} else {
			fmt.Println("Unkwon Errors")
		}

	}
}
