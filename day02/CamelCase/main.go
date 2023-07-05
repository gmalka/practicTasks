package main

import (
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println(camelcase("ABCDFT"))
}

func camelcase(s string) int32 {
	var counter int32 = 1

	if s == "" {
		return 0
	}
	for i := 0; i < len(s); i++ {
		if s[i] >= 'A' && s[i] <= 'Z' {
			counter++
		}
	}

	return counter
}
