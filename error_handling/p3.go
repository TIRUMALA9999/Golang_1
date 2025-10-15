//Write a function safeDivide(a, b int) (int, error) that handles divide-by-zero.

package main
import (
	"fmt"
	"errors"
)

func safeDivide(a,b int) (int,error){
	if b==0{
		return 0,errors.New("divide by zero")
	}
	return a/b,nil
}

func main(){
	i,err:=safeDivide(5,0)
	if err!=nil{
		fmt.Println("error:",err)
	}else{
		fmt.Println(i)
	}
}