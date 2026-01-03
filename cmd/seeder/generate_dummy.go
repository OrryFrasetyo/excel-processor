package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	filename := "large_students.csv"
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	fmt.Println("🚀 Generating 10.000 data...")

	for i := 1; i <= 10000; i++ {
		// data dummy (name & email)
		row := []string{
			"Mahasiswa " + strconv.Itoa(i),
			"mahasiswa" + strconv.Itoa(i) + "@univ.ac.id",
		}
		if err := writer.Write(row); err != nil {
			panic(err)
		}
	}

	fmt.Printf("✅ Selesai! File %s berhasil dibuat.\n", filename)
}