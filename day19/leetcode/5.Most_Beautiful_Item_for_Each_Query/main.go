package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(maximumBeauty([][]int{{10, 1000}}, []int{5}))
}

func maximumBeauty(items [][]int, queries []int) []int {
	sort.Slice(items, func(i, j int) bool {
		return items[i][0] < items[j][0]
	})

	max := 0
	for i := 0; i < len(items); i++ {
		if items[i][1] > max {
			max = items[i][1]
		} else {
			items[i][1] = max
		}
	}

	for k, v := range queries {
		i := sort.Search(len(items), func(i int) bool {
			return items[i][0] > v
		})
		if i != 0 {
			queries[k] = items[i-1][1]
		} else {
			queries[k] = 0
		}
	}

	return queries
}
