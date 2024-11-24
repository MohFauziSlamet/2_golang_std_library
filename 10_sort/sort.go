package main

import (
	"fmt"
	"sort"
)

/*
Package sort
Package sort adalah package yang berisikan utilitas untuk proses pengurutan
Agar data kita bisa diurutkan, kita harus mengimplementasikan kontrak di interface sort.Interface
https://golang.org/pkg/sort/
*/

// * struct user
type User struct {
	Name string
	Age  int
}

// * type alias untuk slice user
type UserSlice []User

// * kontrak interface untuk mendapatkan panjang slice
func (s UserSlice) Len() int {
	return len(s)
}

// * kontrak interface untuk membandingkan dua data , mencari yang terbesar.
func (s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

// * kontrak interface untuk menukar posisi dua data.
func (s UserSlice) Swap(i, j int) {
	//* cara 1
	// temp := s[i]
	// s[i] = s[j]
	// s[j] = temp

	//* cara 2
	//* hal ini bisa dilakukan karena di golang dapat menginisiasi 2 varibel sekaligus.
	//* dan bisa menerima data lebih dari 1 sekaligus. contoh -> ( data, err = GetDb() )
	s[i], s[j] = s[j], s[i]
}

func main() {
	//* membuat slice
	users := []User{
		{"Muhammad", 40},
		{"Azka", 20},
		{"Fitra", 30},
		{"Ramadhan", 10},
	}

	sort.Sort(UserSlice(users))

	println(users)
	fmt.Println(users)
}
