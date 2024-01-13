package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates2(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var ansHead *ListNode = nil
	var ansTail *ListNode = nil
	curr := head
	flag := false

	for curr != nil {
		flag = false
		for curr.Next != nil && curr.Val == curr.Next.Val {
			flag = true
			curr = curr.Next
		}
		if flag {
			curr = curr.Next
			continue
		}
		if ansHead == nil {
			ansHead = curr
			ansTail = ansHead
			curr = curr.Next
		} else {
			ansTail.Next = curr
			ansTail = ansTail.Next
			curr = curr.Next
			// if ansTail != nil && ansTail.Next != nil {
			// 	ansTail.Next = nil
			// }
		}
	}
	if ansTail != nil {
		ansTail.Next = nil
	}

	return ansHead
}

// func main() {
// 	// values := []int{7, 13, 11, 10, 1}
// 	// values := []int{7, 13, 11, 10, 1}
// 	values := []int{1, 1, 1, 1}

// 	nodes := make([]ListNode, len(values))

// 	for ind, val := range values {
// 		nodes[ind] = ListNode{Val: val, Next: nil}
// 	}

// 	for i := 0; i < len(values)-1; i++ {
// 		nodes[i].Next = &nodes[i+1]
// 	}

// 	_ = deleteDuplicates2(&nodes[0])
// }
