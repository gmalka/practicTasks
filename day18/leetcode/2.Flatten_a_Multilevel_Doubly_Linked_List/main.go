package main

func main() {

}

type Node struct {
	Val   int
	Prev  *Node
	Next  *Node
	Child *Node
}

func flatten(root *Node) *Node {
	if root == nil {
		return root
	}
	cur := root

	for cur != nil {
		if cur.Child != nil {
			DoIt(cur)
		}
		cur = cur.Next
	}

	return root
}

func DoIt(node *Node) {
	next := node.Next

	node.Next = node.Child
	node.Child.Prev = node
	node.Child = nil
	node = node.Next
	for node.Next != nil {
		node = node.Next
	}
	node.Next = next
	if next != nil {
		next.Prev = node
	}
}