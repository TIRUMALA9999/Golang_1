/*Fix a bug: don’t continue after error
main() lo first div(10, 2) tarvāta intentionally div(8, 0) call cheyyi.
If error vachindhi ante stop (return) cheyyi; migata lines run kakoodadu.
Print message: "stopping due to error: <err>". */

package main
import (
	"fmt"
	"errors"
)

func div(a,b int) (int,error){
	if b==0{
		return 0,errors.New("division by zero")
	}
	return a/b,nil
}


func main() {
	// First call: valid
	result, err := div(10, 2)
	if err != nil {
		fmt.Println("stopping due to error:", err)
		return // stop execution
	}
	fmt.Println("10 / 2 =", result)

	// Second call: intentional error
	result, err = div(8, 0)
	if err != nil {
		fmt.Println("stopping due to error:", err)
		return // stop execution
	}
	fmt.Println("8 / 0 =", result) // <-- this won’t run
}