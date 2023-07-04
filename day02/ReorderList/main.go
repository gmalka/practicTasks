package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode)  {
    mid, end := head, head

    for end != nil && end.Next != nil {
        mid = mid.Next
        end = end.Next.Next
    }

    pre := mid
	mid = pre.Next
	pre.Next = nil
	for mid != nil {
		c := mid.Next
		mid.Next = pre
		pre = mid
		mid = c
	}

    first := head
    for pre != nil {
        c := first.Next
        first.Next = pre
        pre = c
        first = first.Next
    }
}

func main() {
	task1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7, Next: nil}}}}}}}
	reorderList(task1)

	for task1 != nil {
		fmt.Println(task1.Val)
		task1 = task1.Next
	}
}