//Write a program that creates a file log.txt and appends a new log message every time you run it.

package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	// Open or create file in append mode
	file, err := os.OpenFile("log2.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	// Create the log message
	logMsg := fmt.Sprintf("%s — App started\n",
		time.Now().Format("2006-01-02 15:04:05"))

	// Write the log
	if _, err := file.WriteString(logMsg); err != nil {
		fmt.Println("Error writing log:", err)
		return
	}

	fmt.Println("Log appended successfully ✅")
}
