package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	cur := head

	if head == nil {
		return nil
	}
	length := 1
	for cur.Next != nil {
		length++
		cur = cur.Next
	}

	cur.Next = head
	k = length - (k % length)
	cur = head
	for i := 0; i < k-1; i++ {
		cur = cur.Next
	}
	head = cur.Next
	cur.Next = nil

	return head
}

func main() {
	task1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: nil}}}}}

	task1 = rotateRight(task1, 5)

	for task1 != nil {
		fmt.Println(task1.Val)
		task1 = task1.Next
	}

	fmt.Println("<---------------------------->")

	task2 := &ListNode{Val: 0, Next: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}}
	task2 = rotateRight(task2, 4)

	for task2 != nil {
		fmt.Println(task2.Val)
		task2 = task2.Next
	}
}