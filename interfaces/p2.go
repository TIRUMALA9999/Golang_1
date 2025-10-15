/*Define an interface Vehicle with method Speed() int.

Implement it for:

Car → return 120

Bike → return 80

Write a function printSpeed(v Vehicle) that prints the type and its speed.*/
package main
import "fmt"

// Interface with two methods
type Vehicle interface {
	Speed() int
	Name() string
}

// Car struct
type Car struct {
	Cartype string
}

func (c Car) Speed() int {
	return 180
}
func (c Car) Name() string {
	return c.Cartype
}

// Bike struct
type Bike struct {
	Biketype string
}

func (b Bike) Speed() int {
	return 120
}
func (b Bike) Name() string {
	return b.Biketype
}

// Polymorphic function
func printSpeed(v Vehicle) {
	fmt.Printf("%s --> Speed: %d\n", v.Name(), v.Speed())
}

func main() {
	vals := []Vehicle{Car{"Toyota"}, Bike{"Apache"}}

	for _, val := range vals {
		printSpeed(val)
	}
}
