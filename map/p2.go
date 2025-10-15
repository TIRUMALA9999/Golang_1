//Write a function countFrequency(nums []int) map[int]int → returns frequency of
//  each number in the slice. Example:

/*Input: [1,2,2,3,3,3]
Output: map[1:1 2:2 3:3] */

package main
import "fmt"
func countFrequency(nums []int) map[int]int{
	t:=map[int]int{}
	for _,num:=range nums{
		t[num]++
	}
	return t
}

func main(){
	fmt.Println(countFrequency([]int{1,2,2,3,3,3}))
}