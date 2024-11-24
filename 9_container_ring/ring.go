package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
Package container/ring
Package container/ring adalah implementasi struktur data circular list
Circular list adalah struktur data ring, dimana diakhir element akan kembali ke element awal (HEAD)
https://golang.org/pkg/container/ring/
*/
func main() {
	r := ring.New(5)

	for i := 0; i < r.Len(); i++ {
		r.Value = "Value" + strconv.Itoa(i+1)
		r = r.Next()
	}

	r.Do(func(a any) {
		fmt.Println(a)
	})

}
