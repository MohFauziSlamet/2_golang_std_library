package main

import (
	"encoding/csv"
	"os"
)

func main() {
	writer := csv.NewWriter(os.Stdout)

	writer.Write([]string{"Moh", "Fauzi", "Slamet"})
	writer.Write([]string{"Moh1", "Fauzi1", "Slamet1"})
	writer.Write([]string{"Moh2", "Fauzi2", "Slamet2"})

	writer.Flush() // kita wajib push untuk mengirim semua perubahan

}
