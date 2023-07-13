package main

import (
	"fmt"
)

func main() {
	fmt.Println(countPairs([]int{149, 107, 1, 63, 0, 1, 6867, 1325, 5611, 2581, 39, 89, 46, 18, 12, 20, 22, 234}))
}

func countPairs(deliciousness []int) int {
	count := 0

	m := make(map[int]int, len(deliciousness))
	for _, v := range deliciousness {
		m[v]++
	}

	for n := 0; n < 24; n++ {
		j := 1 << n
		for _, v := range deliciousness {
			if c, ok := m[j-v]; ok {
				if j-v == v {
					count += c - 1
				} else {
					count += c
				}
				
			}
		}
	}

	return (count / 2) % 1000000007
}
