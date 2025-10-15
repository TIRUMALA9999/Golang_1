package main
import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {
	file, _ := os.Create("students.csv")
	defer file.Close()

	writer := csv.NewWriter(file)
	writer.Write([]string{"Name", "Marks"})
	writer.Write([]string{"Alice", "90"})
	writer.Write([]string{"Bob", "85"})
	writer.Flush()

	// Read CSV
	f, _ := os.Open("students.csv")
	defer f.Close()
	reader := csv.NewReader(f)
	records, _ := reader.ReadAll()
	fmt.Println(records)
}
