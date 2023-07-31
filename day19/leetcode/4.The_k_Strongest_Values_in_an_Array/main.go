package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(getStrongest([]int{6, 7, 11, 7, 6, 8}, 5))
}

func getStrongest(arr []int, k int) []int {
	result := make([]int, 0, k)

	sort.Ints(arr)
	l, r := 0, len(arr)-1
	m := (len(arr) - 1) / 2
	for k > 0 {
		la := abs(arr[m] - arr[l])
		ra := abs(arr[m] - arr[r])
		if la > ra || (la == ra && arr[l] >= arr[r]) {
			result = append(result, arr[l])
			k--
			l++
		} else if la < ra || (la == ra && arr[l] <= arr[r]) {
			result = append(result, arr[r])
			k--
			r--
		}
	}

	return result
}

func abs(num int) int {
	if num < 0 {
		return -num
	}

	return num
}
