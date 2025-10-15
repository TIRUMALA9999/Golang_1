//Write a function that reverses a slice of integers in place.

package main
import "fmt"
func reverse(nums []int) []int{
	left:=0
	right:=len(nums)-1
	for left<right{
		nums[left],nums[right]=nums[right],nums[left]
		left++
		right--
	}
	return nums
}

func main() {
    nums := []int{1, 2, 3, 4, 5}
    fmt.Println("Reversed:", reverse(nums))
}