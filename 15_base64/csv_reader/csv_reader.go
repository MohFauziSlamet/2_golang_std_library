package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strings"
)

func main() {

	csvString := "Moh Fauzi Slamet\n" +
		"Muhammad Azkafitra Ramadhan\n" +
		"Barochatul Mauludy"

	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()

		//* untuk menghentikan perulangan.
		//* jika menemui error EOF (end of field)
		if err == io.EOF {
			fmt.Printf("err: %v\n", err)
			break // untuk menghentikan perulangan
		}

		fmt.Printf("record: %v\n", record)
	}
}
