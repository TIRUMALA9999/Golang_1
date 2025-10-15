/* Two Goroutines + Select
Create two goroutines:
One sends "Ping" after 1 second.
Another sends "Pong" after 2 seconds.
Use select in main to print whichever message arrives first.
👉 Expected output:
Ping
Pong */

package main
import (
	"fmt"
	"time"
)
func main(){
	ch1:=make(chan string)
	ch2:=make(chan string)
	go func(){
		time.Sleep(1 * time.Second)
		ch1 <- "Ping"
	}()
	go func(){
		time.Sleep(2 * time.Second)
		ch2 <- "Pong"
	}()

	for i:=0;i<2;i++{
		select{
	    case msg1:= <- ch1:
		    fmt.Println(msg1)
	    case msg2:= <- ch2:
		    fmt.Println(msg2)
		
		}
	}
	
}

/* Two separate channels (ch1, ch2)
Each goroutine ki okka dedicated channel untundi.
Easy ga identify cheyyagalav → Ping ochindaa, Pong ochindaa.
select lo separate cases rayadam simple. */