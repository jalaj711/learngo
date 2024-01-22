package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}
	count := 1
	var prev *ListNode
	curr := head
	for count < left {
		prev = curr
		curr = curr.Next
		count++
	}
	leftnode := prev
	prev = curr
	curr = curr.Next
	count++

	var temp *ListNode

	for count < right {
		temp = curr.Next
		curr.Next = prev
		prev = curr
		curr = temp
		count++
	}

	if leftnode == nil {
		head.Next.Next = curr.Next
		curr.Next = head
		return curr
	}

	leftnode.Next.Next = curr.Next
	curr.Next = prev
	leftnode.Next = curr

	return head
}

// func main() {
// 	// values := []int{7, 13, 11, 10, 1}
// 	// values := []int{7, 13, 11, 10, 1}
// 	values := []int{3, 5}

// 	nodes := make([]ListNode, len(values))

// 	for ind, val := range values {
// 		nodes[ind] = ListNode{Val: val, Next: nil}
// 	}

// 	for i := 0; i < len(values)-1; i++ {
// 		nodes[i].Next = &nodes[i+1]
// 	}

// 	_ = reverseBetween(&nodes[0], 1, 2)
// }
