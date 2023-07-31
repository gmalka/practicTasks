package main

import (
	"fmt"
	"sort"
)

func main() {

	fmt.Println(countFairPairs([]int{0, 1, 7, 4, 4, 5}, 3, 6))

}

func countFairPairs(nums []int, lower int, upper int) int64 {
	sort.Ints(nums)

	sum := 0
	l, ll := len(nums)-1, len(nums)-1

	i := 0
	for i != ll {
		if l > ll {
			l = ll
		}
		if nums[i]+nums[ll] > upper {
			ll--
		} else if i >= l || nums[i]+nums[l] < lower {
			// fmt.Println(i, l, nums[i] + nums[l])
			i++
			sum += ll - l
			l++
			// fmt.Println(sum)
		} else {
			l--
		}
	}

	return int64(sum)
}
