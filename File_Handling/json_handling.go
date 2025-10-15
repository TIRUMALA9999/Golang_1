package main
import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Marks int
}

func main() {
	students := []Student{
		{"Alice", 90},
		{"Bob", 85},
	}

	// Write JSON
	data, _ := json.MarshalIndent(students, "", "  ")
	os.WriteFile("students.json", data, 0644)

	// Read JSON
	fileData, _ := os.ReadFile("students.json")
	var decoded []Student
	json.Unmarshal(fileData, &decoded)

	fmt.Println("Decoded:", decoded)
}


/* First manam Student ani struct create chesam.

Then slice lo Alice, Bob ni add chesam.

json.MarshalIndent tho JSON format ki convert chesam.

os.WriteFile tho file lo rayadam.

Tarvata os.ReadFile tho file chadavadam.

json.Unmarshal tho malli Go struct ki convert chesam.

Finally print chesam ✅ */