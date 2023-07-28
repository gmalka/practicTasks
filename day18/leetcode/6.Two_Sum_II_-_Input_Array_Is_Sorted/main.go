package main

import (
	"fmt"
)

func main() {
	fmt.Println(twoSum([]int{-1, 0}, -1))
}

func twoSum(numbers []int, target int) []int {
	l, r := 0, len(numbers)-1
	for {
		for numbers[l]+numbers[r] > target {
			r--
		}
		if numbers[l]+numbers[r] == target && r != l {
			return []int{l + 1, r + 1}
		}

		fmt.Println(l, r)
		for numbers[l]+numbers[r] < target {
			l++
		}
		if numbers[l]+numbers[r] == target && r != l {
			return []int{l + 1, r + 1}
		}
	}
}
