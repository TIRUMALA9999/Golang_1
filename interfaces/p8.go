/*🟡 Problem 8: Type Switch

Write a function:

func checkType(x interface{}) {
    // should print the type category
}

Requirements:

If x is an int → print "Integer"

If x is a string → print "String"

Otherwise → print "Something else"

Example Usage:
checkType(42)
checkType("Teja")
checkType(true)

Expected Output:
Integer
String
Something else*/

package main
import "fmt"

func checkType(x interface{}) {
	switch v := x.(type) {
	case int:
		fmt.Println("Integer")
	case string:
		fmt.Println("String")
	default:
		fmt.Printf("Something else (type: %T, value: %v)\n", v, v)
	}
}

func main() {
	checkType(42)       // Integer
	checkType("Teja")   // String
	checkType(true)     // Something else
}
