package main

func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	cnt := 1
	curr := list1
	for cnt < a {
		curr = curr.Next
		cnt++
	}
	temp := curr.Next
	curr.Next = list2
	curr = temp

	for cnt < b {
		curr = curr.Next
		cnt++
	}

	for list2.Next != nil {
		list2 = list2.Next
	}

	list2.Next = curr.Next
	return list1
}
