package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
Package io :
IO atau singkatan dari Input Output, merupakan fitur di Golang yang
digunakan sebagai standard untuk proses Input Output.
Di Golang, semua mekanisme input output pasti mengikuti standard package io.
https://pkg.go.dev/io
*/

func main() {
	fmt.Println()
	fmt.Println("==============================")
	fmt.Println()

	input := strings.NewReader("muhammad azkafitra ramadhan \nbarochatul mauludy")

	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()

		//* stop loop
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))

	}

	fmt.Println()
	fmt.Println("==============================")
	fmt.Println()

	writer := bufio.NewWriter(os.Stdout)

	writer.WriteString("Hello world \n")
	writer.WriteString("Selamat belajar")

	writer.Flush()

	fmt.Println()
	fmt.Println()
	fmt.Println("==============================")
	fmt.Println()
}
