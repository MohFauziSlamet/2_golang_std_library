package main

import (
	"fmt"
	"time"
)

/*
Duration
Saat menggunakan tipe data waktu, kadang kita butuh data berupa durasi
Package tipe memiliki type Duration, yang sebenarnya adalah alias untuk int64
Namun terdapat banyak method yang bisa kita gunakan untuk memanipulasi Duration
*/
func main() {
	d := (time.Second * 100)
	d2 := (time.Millisecond * 10)

	fmt.Printf("d: %d\n", d)
	fmt.Printf("d2: %d\n", d2)
	fmt.Printf("(d - d2): %d\n", (d - d2))
}
