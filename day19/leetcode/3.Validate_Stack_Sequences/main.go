package main

import "fmt"

func main() {
	fmt.Println(validateStackSequences([]int{1, 2, 3, 0}, []int{2, 1, 3, 0}))
}

func validateStackSequences(pushed []int, popped []int) bool {
	push, pop := 0, 0
	p := make([]int, 0, len(pushed))

	for push < len(pushed) {
		if len(p) > 0 && p[len(p)-1] == popped[pop] {
			pop++
			p = p[:len(p)-1]
			continue
		} else if popped[pop] == pushed[push] {
			pop++
		} else {
			p = append(p, pushed[push])
		}
		push++
	}

	for i := len(p) - 1; i >= 0; i-- {
		if pop > len(popped) || p[i] != popped[pop] {
			return false
		}

		pop++
	}

	return true
}
