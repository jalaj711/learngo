package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	prev := head
	curr := head.Next

	for curr != nil {
		if curr.Val == prev.Val {
			curr = curr.Next
			continue
		}
		prev.Next = curr
		prev = curr
		curr = curr.Next
	}
	prev.Next = curr

	return head
}

// func main() {
// 	// values := []int{7, 13, 11, 10, 1}
// 	// values := []int{7, 13, 11, 10, 1}
// 	values := []int{3, 3} //, 4, 4, 4, 5, 6, 7, 7, 7}

// 	nodes := make([]ListNode, len(values))

// 	for ind, val := range values {
// 		nodes[ind] = ListNode{Val: val, Next: nil}
// 	}

// 	for i := 0; i < len(values)-1; i++ {
// 		nodes[i].Next = &nodes[i+1]
// 	}

// 	_ = deleteDuplicates(&nodes[0])
// }
