package main

import "fmt"

func main() {
	fmt.Println(maxScore([]int{1,2,3,4,5,6,1}, 3))
}

func maxScore(cardPoints []int, k int) int {
	cur, max := 0, 0
	for _, v := range cardPoints[:k] {
		cur += v
		max = cur
	}
	l, r := k - 1, len(cardPoints) - 1
	for l >= 0 {
		cur -= cardPoints[l]
		cur += cardPoints[r]
		l--
		r--
		if cur > max {
			max = cur
		}
	}

	return max
}