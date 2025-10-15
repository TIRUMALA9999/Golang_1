/*Define an interface Notifier with method Notify() string.

Implement it for:

EmailNotifier → returns "Sending Email..."

SMSNotifier → returns "Sending SMS..."

Write a function sendNotification(n Notifier) that prints the result of Notify(). */

package main
import "fmt"

type Notifier interface{
	Notify() string
}

type EmailNotifier struct{}

func (e EmailNotifier) Notify() string{
	return "Sending Email..."
}

type SMSNotifier struct{}

func (s SMSNotifier) Notify() string{
	return "Sending SMS..."
}

func sendNotification(n Notifier) string{
	return n.Notify()
}

func main(){
	e1:=EmailNotifier{}
	s1:=SMSNotifier{}
	fmt.Println(sendNotification(e1))
	fmt.Println(sendNotification(s1))
}
