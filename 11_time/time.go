package main

import (
	"fmt"
	"time"
)

/*
Package time

Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
https://golang.org/pkg/time/


Beberapa Function di Package time
Function					Kegunaan
time.Now()					Untuk mendapatkan waktu saat ini
time.Date(...)				Untuk membuat waktu
time.Parse(layout, string)	Untuk memparsing waktu dari string
*/

func main() {
	date := time.Now()
	fmt.Printf("time.Local(): %v\n", date.Local())
	fmt.Println(date.Local())
	fmt.Println(date.Clock())
	fmt.Println(date.Day())
	fmt.Println(date.GoString())
	fmt.Println(date.Hour())
	fmt.Println(date.Location())
	fmt.Println(date.ISOWeek())
	fmt.Println(date.IsZero())
	fmt.Println(date.Minute())
	fmt.Println(date.Month())

	utc := time.Date(2024, time.March, 10, 0, 0, 0, 0, time.UTC)

	fmt.Printf("utc: %v\n", utc)
	fmt.Printf("utc: %v\n", utc.Local())

	formater := "2006-01-01 15:04:01"

	value := "2010-10-10 10:10:10"

	t, err := time.Parse(formater, value)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(t)

	}

}
