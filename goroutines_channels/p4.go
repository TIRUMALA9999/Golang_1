//Write a program where two goroutines send messages into a channel, and main prints whichever
//  comes first (using select).

package main
import "fmt"
func main(){
	ch1:=make(chan string)
	ch2:=make(chan string)
	go func(){ch1 <- " Hello"}()
	go func(){ch2 <- "Teja"}()

	select{
	case msg1:=<-ch1:
		fmt.Println(msg1)
	case  msg2:=<-ch2:
		fmt.Println(msg2)
	}
}