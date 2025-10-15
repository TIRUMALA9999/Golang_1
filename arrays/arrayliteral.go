package main
import "fmt"
func main(){
	nums := [5]int{1, 2, 3, 4, 5}    //initialising and declaring at the same time. this is 
	                                 //array literal
	fmt.Println(nums) // [1 2 3 4 5]

	nums1 := [...]int{10, 20, 30}  //auto length inference
	fmt.Println(len(nums1)) // [10 20 30]

	nums2 := [5]int{1, 2}   //automatically assign zeroes to unspecifies numbers.
	fmt.Println(nums2) // [1 2 0 0 0]

}