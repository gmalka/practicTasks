package main

import (
	"fmt"
)

func main() {
	result := maximumXOR([]int{3, 2, 4, 6})

	fmt.Println(result)
}

func maximumXOR(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] |= nums[i]
	}

	return nums[0]
}
