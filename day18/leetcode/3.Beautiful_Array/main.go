package main

func main() {

}

func beautifulArray(n int) []int {
	result := []int{1}

	for len(result) < n {
		newResult := make([]int, 0, n)
		for _, v := range result {
			if v*2-1 <= n {
				newResult = append(newResult, v*2-1)
			}
		}
		for _, v := range result {
			if v*2 <= n {
				newResult = append(newResult, v*2)
			}
		}
		result = newResult
	}

	return result
}
