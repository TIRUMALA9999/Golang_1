//Read a CSV of numbers and compute their sum.

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Write CSV
	file, err := os.Create("numbers1.csv")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	writer := csv.NewWriter(file)

	if err := writer.Write([]string{"1","2","3","4","5","6","7","8","9","10"}); err != nil {
		fmt.Println("Error writing CSV:", err)
		return
	}
	writer.Flush()
	if err := writer.Error(); err != nil {
		fmt.Println("Writer error:", err)
		return
	}
	// ensure data hits disk
	_ = file.Sync()
	_ = file.Close()

	// Read CSV
	f, err := os.Open("numbers1.csv")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
		return
	}

	sum := 0
	for _, row := range records {
		for _, value := range row {
			num, err := strconv.Atoi(value)
			if err != nil {
				fmt.Println("Conversion error:", err)
				continue
			}
			sum += num
		}
	}

	fmt.Println("✅ Numbers in file:", records)
	fmt.Println("🧮 Sum of numbers:", sum)
}
