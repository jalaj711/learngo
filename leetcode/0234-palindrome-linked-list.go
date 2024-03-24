package main

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	fast, slow, next := head, head, head.Next
	var prev *ListNode

	for slow.Next != nil && fast.Next != nil && (fast.Next).Next != nil {
		fast = fast.Next.Next

		slow.Next = prev
		prev = slow
		slow = next
		next = slow.Next

	}
	slow.Next = prev
	if fast.Next == nil && fast != head {
		// if fast.Next == nil, then there were odd number of elements in the list
		// and the slow pointer is currently at the middle element
		slow = prev
	}

	for slow != nil && next != nil {
		if slow.Val != next.Val {
			return false
		}
		slow = slow.Next
		next = next.Next
	}

	return true
}

//
//func main() {
//	g := ListNode{1, nil}
//	f := ListNode{2, &g}
//	e := ListNode{3, &f}
//	d := ListNode{4, &e}
//	c := ListNode{3, &d}
//	b := ListNode{2, &c}
//	a := ListNode{1, &b}
//
//	isPalindrome(&a)
//}
