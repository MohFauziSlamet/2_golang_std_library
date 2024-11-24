package main

import (
	"fmt"
	"reflect"
)

/*
Package reflect
Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa
melihat struktur kode kita pada saat aplikasi sedang berjalan.
Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect.
Fitur ini mungkin tidak bisa dibahas secara lengkap dalam satu video,
Anda bisa eksplorasi package reflec ini secara otodidak
Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
https://golang.org/pkg/reflect/

*/

type Sample struct {
	Name string
}

type Person struct {
	Name, Address, Email string
}

func readField(value any) {
	valueType := reflect.TypeOf(value)

	fmt.Println("Value Type", valueType)

	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println("Value Field : ", valueField, " with type ", valueType)

	}
}

func main() {
	readField(Sample{"Fitra"})
	readField(Person{"Azka", "Sumberpucung", "azka@gmail.com"})
}
