//Define a Book struct with fields: title, author, price. Create 2 books and print details.

package main
import "fmt"
 
type Book struct{
	title string
	author string
	price string
}

func main(){
	b1:=Book{title:"Teja",author:"atj",price:"$10"}
	b2:=Book{title:"Jaja",author:"kpg",price:"$11"}
	fmt.Println(b1.title," ",b1.author," ",b1.price)
	fmt.Println(b2.title," ",b2.author," ",b2.price)
}