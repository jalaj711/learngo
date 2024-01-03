package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var sorted *ListNode = nil
	var prev *ListNode = nil

	current := head

	for sorted != head {
		for current.Next != sorted {
			if current.Val > current.Next.Val {
				if current == head {
					prev = current.Next
					current.Next = prev.Next
					prev.Next = current
					head = prev
				} else {
					prev.Next = current.Next
					current.Next = prev.Next.Next
					prev = prev.Next
					prev.Next = current
				}
			} else {
				if current != head {
					prev = prev.Next
				} else {
					prev = head
				}
				current = current.Next
			}
		}
		sorted = current
		current = head
	}

	return sorted
}

// func main() {
// 	// values := []int{7, 13, 11, 10, 1}
// 	// values := []int{7, 13, 11, 10, 1}
// 	values := []int{16, 13, 11, 10, 1}

// 	nodes := make([]ListNode, len(values))

// 	for ind, val := range values {
// 		nodes[ind] = ListNode{Val: val, Next: nil}
// 	}

// 	for i := 0; i < len(values)-1; i++ {
// 		nodes[i].Next = &nodes[i+1]
// 	}

// 	_ = sortList(&nodes[0])
// }
