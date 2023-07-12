package main

import "fmt"

func main() {
	fmt.Println(generateParenthesis(1))
}

func generateParenthesis(n int) []string {
    return recursive([]string{"()"}, n-1, 1)
}

func recursive(str []string, n, offset int) []string {
	if n == 0 {
		return str
	}
	
	s := "()"
	result := make([]string, 0, 10)
	for ; offset <= len(str[0]); offset++ {
		in := str[0][:offset] + s + str[0][offset:]
		result = append(result, recursive([]string{in}, n-1, offset + 1)...)
	}

	return result
}

// () -> (()), ()() -> ((())), (()()), (())(), ()(()), ()()()