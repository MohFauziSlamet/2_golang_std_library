package main

import (
	"fmt"
	"path"
	"path/filepath"
)

/*
Package path :
Package path digunakan untuk memanipulasi data path seperti path di URL atau path di File System.
Secara default Package path menggunakan slash sebagai karakter path nya, oleh karena itu cocok untuk data URL
https://pkg.go.dev/path
Namun jika ingin menggunakan untuk memanipulasi path di File System,
karena Windows menggunakan backslash, maka khusus untuk File System, perlu menggunakan pacakge path/filepath
https://pkg.go.dev/path/filepath
*/

func main() {
	fmt.Println(path.Dir("hallo/hello/world.go")) // folder path file
	fmt.Println(path.Base("hello/world.go"))      // nama file
	fmt.Println(path.Ext("hello/world.go"))       // extention file
	fmt.Println(path.Join("hello", "world", "main.go"))

	fmt.Println()
	fmt.Println("==============================")
	fmt.Println()

	fmt.Println(filepath.Dir("hallo/hello/world.go")) // folder path file
	fmt.Println(filepath.Base("hello/world.go"))      // nama file
	fmt.Println(filepath.Ext("hello/world.go"))       // extention file
	fmt.Println(filepath.Join("hello", "world", "main.go"))
}
