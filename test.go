package main

import (
	// "strconv"
	"fmt"
)

func main () {
	var nums []int = []int{1,2,3}
	for i := 0; i< len(nums); i++ {
		fmt.Println(nums[i])
		if(nums[i] == 3) {
			nums = append(nums, 4)
		}
	}
	fmt.Printf("%v", nums)
}
