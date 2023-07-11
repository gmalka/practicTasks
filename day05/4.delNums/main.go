package main

import "fmt"

func main() {
	fmt.Println(DelK([]int{1, 2, 3, 3, 4}, 2))
	fmt.Println(DelK([]int{1, 2, 3, 3, 4}, 1))
	fmt.Println(DelK([]int{1, 2, 3, 3, 4, 3}, 2))
		fmt.Println(DelK([]int{1, 1, 2, 3, 3, 4, 3, 1}, 3))
}

func DelK(arr []int, k int) []int {
	count := make(map[int]int, len(arr))
	for i := 0; i < len(arr); i++ {
		count[arr[i]]++
	}

	l, r := 0, 0

	for ; r < len(arr); r++ {
		if v, ok := count[arr[r]]; ok && v < k {
			arr[l] = arr[r]
			l++
		}
	}

	return arr[:l]
}