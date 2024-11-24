package main

import (
	"fmt"
	"slices"
)

/*
Package slices
Di Golang versi terbaru, terdapat fitur bernama Generic,
fitur ini akan kita bahas khusus dikelas Golang Generic.
Fitur Generic ini membuat kita bisa membuat parameter
dengan tipe yang bisa berubah-ubah, tanpa harus menggunakan interface kosong / any.
Salah satu package yang menggunakan fitur Generic ini adalah package slices.
Package slices ini digunakan untuk memanipulasi data di slice.
https://pkg.go.dev/slices

*/

func main() {
	names := []string{"Jonh", "Paul", "George", "Raana", "Ringo", "Riingo"}
	values := []int{10, 11, 12, 13, 14}

	fmt.Printf("slices.Min(names): %v\n", slices.Min(names))
	fmt.Printf("slices.Max(names): %v\n", slices.Max(names))
	fmt.Printf("slices.Min(values): %v\n", slices.Min(values))
	fmt.Printf("slices.Max(values): %v\n", slices.Max(values))
	slices.Contains(names, "Paul")
	slices.Index(names, "Paul")

}
