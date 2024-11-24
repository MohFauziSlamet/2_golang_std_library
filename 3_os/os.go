package main

/*
	Package os :
	Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os.
	Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen
	(bisa digunakan  disemua sistem operasi).
	https://golang.org/pkg/os/
*/
import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	fmt.Println(args)

	for i, v := range args {
		fmt.Println(i, "=>", v)
	}

	name, err := os.Hostname()
	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err)

	}
}
