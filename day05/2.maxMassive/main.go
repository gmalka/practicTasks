package main

import "fmt"

func main() {
	fmt.Println(maxSubarray([]int{10, 10, 10, 3, 11, 11, 12}))
	fmt.Println(maxSubarray([]int{}))
	fmt.Println(maxSubarray([]int{1}))
	fmt.Println(maxSubarray([]int{2, 1}))
	fmt.Println(maxSubarray([]int{10, 10, 10, 5, 5, 3}))
}

func maxSubarray(arr []int) []int {
	maxL, maxR := 0, 0
	l, r := 0, 0
	sum := 0
	max := 0

	if len(arr) == 0 {
		return nil
	}
	for ; r < len(arr); r++ {
		maxR = r
		if arr[l] == arr[r] {
			sum += arr[l]
		} else {
			break
		}
	}


	pre := r
	for ; r < len(arr); r++ {
		if arr[pre] == arr[r] {
			sum += arr[r]
		} else {
			if sum > max {
				max = sum
				maxL, maxR = l, r
			}
			for arr[l] != arr[pre] {
				sum -= arr[l]
				l++
			}
			sum += arr[r]
		}
		pre = r
	}

	if sum > max {
		return arr[l:r]
	}
	return arr[maxL:maxR]
}

// 10 10 2 3 4