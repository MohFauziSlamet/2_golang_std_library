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
	Name string `required:"true" max:"10"`
}

type Person struct {
	Name    string `required:"true" max:"10"`
	Address string `required:"true" max:"10"`
	Email   string `required:"true" max:"10"`
}

func readField(value any) {
	valueType := reflect.TypeOf(value)

	fmt.Println("Value Type", valueType)

	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		fmt.Println("Value Field : ", valueField, " with type ", valueType)

		fmt.Printf("valueField.Tag.Get(\"required\"): %v\n", valueField.Tag.Get("required"))
		fmt.Printf("valueField.Tag.Get(\"max\"): %v\n", valueField.Tag.Get("max"))
		fmt.Printf("valueField.Tag.Get(\"min\"): %v\n", valueField.Tag.Get("min"))

	}
}

func IsValid(value any) bool {
	res := true

	//* ambil type value
	typeValue := reflect.TypeOf(value)

	//* looping data value.
	//* check tag , apakah ada yang true didalam struct.
	for i := 0; i < typeValue.NumField(); i++ {

		//* ambil setiap field yang ada di struct
		structField := typeValue.Field(i)

		//* check apakah ada yang true didalam field
		if structField.Tag.Get("required") == "true" {
			//* jika ada , do this
			dataInterfaceValue := reflect.ValueOf(value).Field(i).Interface()

			//* log isi data
			fmt.Println("dataInterfaceValue: ", i, dataInterfaceValue)

			res = dataInterfaceValue != ""
			if !res {
				fmt.Printf("res 1: %v\n", res)
				return res
			}
		}

	}
	fmt.Printf("res 2: %v\n", res)
	return res

}

func main() {
	readField(Sample{"Fitra"})
	readField(Person{"Azka", "Sumberpucung", "azka@gmail.com"})

	person := Person{Name: "nama", Address: "address", Email: "email"}

	fmt.Printf("IsValid(person): %v\n", IsValid(person))
}
