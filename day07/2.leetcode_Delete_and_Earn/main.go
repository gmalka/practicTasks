package main

import "fmt"

func main() {
	fmt.Println(deleteAndEarn([]int{3,4,2}))
}

func deleteAndEarn(nums []int) int {
    max := 0

	for i := 0; i < len(nums); i++ {
		if max < nums[i] {
			max = nums[i]
		}
	}

	dp := make([]int, max + 1)

	for i := 0; i < len(nums); i++ {
		dp[nums[i]] += nums[i]
	}

	for i := 2; i < len(dp); i++ {
		if dp[i] + dp[i - 2] > dp[i - 1] {
			dp[i] = dp[i] + dp[i - 2]
		} else {
			dp[i] = dp[i - 1]
		}
	}

	return dp[len(dp) - 1]
}