package main

import (
	"fmt"
)

func main() {
	k := make([]int, 4, 4)
	for i := 0; i < 4; i++ {
		k[i] = i
	}
	k = append(k, []int{5, 6, 7, 8, 9, 10})
	fmt.Println(k, len(k), cap(k))

	s := make([]string, 0, 4)
	s = append(s, []string{"hi", "i", "am", "from", "Kazan"})
	s = append(s, []string{",and", "i", "learning", "golang", ";)"})
	fmt.Println(s, len(s), cap(s))
}

func append[T any](origin, appended []T) []T {
	if cap(origin)-len(origin) < len(appended) {
		l := cap(origin)
		for l-len(origin) < len(appended) {
			l *= 2
		}
		newArr := make([]T, l)
		copy(newArr, origin)
		origin = newArr[:len(origin)]
	}

	originLen := len(origin)
	origin = origin[:originLen+len(appended)]
	copy(origin[originLen:originLen+len(appended)], appended)

	return origin
}
