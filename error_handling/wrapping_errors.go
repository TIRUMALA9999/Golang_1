/*👉 %w lets you wrap an error with context, while keeping the original inside.
This means callers can still check with errors.Is or errors.As. */

package main
import (
	"fmt"
	"os"
	"errors"
)

func loadConfig(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		// add context + wrap original
		return nil, fmt.Errorf("reading config %q: %w", path, err)
	}
	return data, nil
}

func main(){
	_, err := loadConfig("missing.yml")
    if err != nil {
	    fmt.Println("failed:", err)
	    if errors.Is(err, os.ErrNotExist) {
		    fmt.Println("reason: file does not exist")
	}
}
}
