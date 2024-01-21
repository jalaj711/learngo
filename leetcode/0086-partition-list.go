package main

func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	var p1head *ListNode = nil
	var p2head *ListNode = nil
	var p1tail *ListNode = nil
	var p2tail *ListNode = nil
	curr := head

	for curr != nil {
		if curr.Val < x {
			if p1head == nil {
				p1head = curr
				p1tail = curr
			} else {
				p1tail.Next = curr
				p1tail = curr
			}
		} else {
			if p2head == nil {
				p2head = curr
				p2tail = curr
			} else {
				p2tail.Next = curr
				p2tail = curr
			}
		}
		curr = curr.Next
	}

	if p2tail != nil {
		p2tail.Next = nil
	}

	if p1head == nil || p1tail == nil {
		return p2head
	}

	p1tail.Next = p2head
	return p1head
}

// func main() {
// 	values := []int{1, 1}
// 	x := 0
// 	// values := []int{-1, -1}
// 	// x := 0
// 	// values := []int{1, 4, 3, 2, 5, 2}
// 	// x:= 3

// 	nodes := make([]ListNode, len(values))

// 	for ind, val := range values {
// 		nodes[ind] = ListNode{Val: val, Next: nil}
// 	}

// 	for i := 0; i < len(values)-1; i++ {
// 		nodes[i].Next = &nodes[i+1]
// 	}

// 	_ = partition(&nodes[0], x)
// }
