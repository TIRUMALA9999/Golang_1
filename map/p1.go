package main
import "fmt"
func main(){
	t:=map[string]int{
		"delhi": 45677878,
		"mumbai":35654745,
		"goa":457880654,
	}
	fmt.Println(t)
	t["mumbai"]=34654754
	fmt.Println(t)
	delete(t,"mumbai")
	fmt.Println(t)
}