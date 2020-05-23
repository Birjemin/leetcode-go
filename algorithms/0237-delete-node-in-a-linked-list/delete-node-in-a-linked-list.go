package delete_node_in_a_linked_list

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNode(node *ListNode) {
	for node != nil {
		node.Val = node.Next.Val
		if node.Next.Next == nil {
			node.Next = nil
		}
		node = node.Next
	}
}

func deleteNode1(node *ListNode) {
	node.Val, node.Next = node.Next.Val, node.Next.Next
}
