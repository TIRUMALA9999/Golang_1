/* Create a slice of type []interface{}.
Store different values in it:
25 (int)
"Teja" (string)
true (bool)
Guitar{} (struct from before)
Iterate through the slice and print both type and value (like Problem 5).*/

package main
import "fmt"
type Guitar struct{}
func main(){
	var x []interface{}
    x = append(x, 25, "Teja", true, Guitar{})
	for _,val:=range x{
		fmt.Printf("Type: %T,Value: %v\n",val,val)
	}
}

/*Rule of Thumb

If you don’t know how many values → use append.

If you know exact length upfront → use make and assign with indices.*/

/*👉 So, slices are dynamic ✅ … but you can’t directly assign to an index that doesn’t exist yet.
 You must either preallocate or grow with append.*/