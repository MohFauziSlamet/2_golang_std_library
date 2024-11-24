package main

import (
	"encoding/base64"
	"fmt"
)

/*
Package encoding
Golang menyediakan package encoding untuk melakukan encode dan decode
https://pkg.go.dev/encoding
Golang menyediakan berbagai macam algoritma untuk encoding, contoh yang populer adalah base64, csv dan json

*/

func main() {
	value := "Muhammad Azkafitra Ramadhan"
	fmt.Printf("value: %v\n", value)

	//* encoded
	encoded := base64.StdEncoding.EncodeToString([]byte(value))

	fmt.Printf("encoded: %v\n", encoded)

	//* byteDecoded
	byteDecoded, err := base64.StdEncoding.DecodeString(encoded)

	if err != nil {
		fmt.Printf("err: %v\n", err)
	} else {

		fmt.Printf("decoded: %v\n", string(byteDecoded))
		fmt.Printf("decoded: %v\n", byteDecoded)
	}

}
