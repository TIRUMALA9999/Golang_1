//Write a function sqrt(n float64) (float64, error) that returns an error if n < 0.

package main
import (
	"fmt"
	"errors"
	"math"
)

func sqrt(n float64) (float64,error){
	if n<0{
		return 0,errors.New("You given negative number")
	}
	return math.Sqrt(n),nil
}

func main(){
	i,err:=sqrt(5)
	if err!=nil{
		fmt.Println("error is:",err)
	}else{
		fmt.Println(i)
	}
}
