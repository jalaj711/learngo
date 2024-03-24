package main

func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	slow := head
	fast := head
	// find the mid-point
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// reverse the second half
	var prev *ListNode = nil
	next := slow.Next
	slow.Next = nil
	slow = next
	for next != nil {
		next = slow.Next
		slow.Next = prev
		prev = slow
		slow = next
	}

	// merge the two lists
	slow = head
	next = prev
	for slow != nil && next != nil {
		prev = slow.Next
		slow.Next = next
		slow = next
		next = prev
	}
	return
}

func main() {
	g := ListNode{7, nil}
	f := ListNode{6, &g}
	e := ListNode{5, &f}
	d := ListNode{4, &e}
	c := ListNode{3, &d}
	b := ListNode{2, &c}
	a := ListNode{1, &b}

	reorderList(&a)
}
