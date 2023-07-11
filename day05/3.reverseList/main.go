package main

import "fmt"

type Node struct {
	Val  int
	Next *Node
}

func main() {
	list := &Node{1, &Node{2, &Node{3, &Node{4, &Node{5, &Node{6, &Node{7, nil}}}}}}}
	for head := list; head != nil; head = head.Next {
		fmt.Print(head.Val, " ")
	}
	fmt.Println()

	list = ReverseList(list)
	for list != nil {
		fmt.Print(list.Val, " ")
		list = list.Next
	}

	fmt.Println()
	list = &Node{1, nil}
	list = ReverseList(list)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}

	fmt.Println()
	list = nil
	list = ReverseList(list)
	for list != nil {
		fmt.Println(list.Val)
		list = list.Next
	}
}

func ReverseList(head *Node) *Node {
	if head == nil {
		return nil
	}

	var pre *Node = nil
	first := head

	for head != nil {
		head = first.Next
		first.Next = pre
		pre = first
		first = head
	}

	return pre
}
