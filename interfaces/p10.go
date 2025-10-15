/*Polymorphism in Action
👉 Define an interface Payment with method:
Pay(amount int) string
Task:
Implement Payment for two types:
CreditCard → returns "Paid 100 using Credit Card"
PayPal → returns "Paid 100 using PayPal"
Write a function checkout(p Payment, amount int) that prints the result of Pay(amount).
In main(), call checkout with both payment methods.
Expected Output:
Paid 100 using Credit Card
Paid 100 using PayPal*/

package main
import "fmt"

type Payment interface{
	Pay(amount int) string
}

type CreditCard struct{}
func (c CreditCard) Pay(amount int) string{
	return fmt.Sprintf("Paid %d using Creditcard",amount)
}

type PayPal struct{}
func (p1 PayPal) Pay(amount int) string{
	return fmt.Sprintf("Paid %d using PayPal",amount)	
}

func checkout(p Payment, amount int){
	fmt.Println(p.Pay(amount))
}

func main(){
	cc := CreditCard{}
	pp := PayPal{}

	checkout(cc, 100) // Paid 100 using Credit Card
	checkout(pp, 110) // Paid 100 using PayPal
}