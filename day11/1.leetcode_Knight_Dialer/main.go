package main

import "fmt"

func knightDialer(n int) int {
	result := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	m := map[int][]int{
		1: {6, 8},
		2: {7, 9},
		3: {4, 8},
		4: {3, 9, 0},
		5: {},
		6: {1, 7, 0},
		7: {2, 6},
		8: {3, 1},
		9: {2, 4},
		0: {6, 4},
	}

	for i := 1; i < n; i++ {
		newResult := make([]int, 10)
		for j := 0; j < len(result); j++ {
			for _, v := range m[j] {
				newResult[v] += result[j]
				newResult[v] %= 1000000007
			}
		}
		result = newResult
	}

	sum := 0
	for _, v := range result {
		sum += v
		sum %= 1000000007
	}

	return sum
}

func main() {
	fmt.Println(knightDialer(3131))
}
