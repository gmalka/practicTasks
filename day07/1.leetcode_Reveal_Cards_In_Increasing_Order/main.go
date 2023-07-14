package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(deckRevealedIncreasing([]int{17,13,11,2,3,5,7}))
}

type list struct {
	begin, end *node
}

type node struct {
	next *node
	val int
}

func (l *list) add(val int) {
	if l.begin == nil {
		l.begin = &node{val: val}
		l.end = l.begin
		return
	}

	l.end.next = &node{val: val}
	l.end = l.end.next
}

func (l *list) pop() (int, bool) {
	if l.begin != nil {
		val := l.begin.val
		l.begin.next, l.begin = nil, l.begin.next
		if l.begin == nil {
			l.end = nil
		}
		return val, true
	}

	return 0, false
}

func deckRevealedIncreasing(deck []int) []int {
	sort.Ints(deck)
	l := &list{}

	result := make([]int, len(deck))
	for i := 0; i < len(deck); i++ {
		l.add(i)
	}

	for i := 0; i < len(deck); i++ {
		val, _ := l.pop()

		result[val] = deck[i]

		if i != i - 1 {
            val, _ := l.pop()
			l.add(val)
        }
	}

	return result
}