package main

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	current := (*head).Next
	nHead := &Node{Val: (*head).Val, Next: nil, Random: nil}
	nCurrent := nHead

	oldToNew := make(map[*Node]*Node)
	oldToNew[head] = nCurrent

	for current != nil {
		(*nCurrent).Next = &Node{Val: current.Val, Next: current.Next, Random: nil}
		oldToNew[current] = nCurrent.Next

		current = (*current).Next
		nCurrent = (*nCurrent).Next
	}

	current = head
	nCurrent = nHead

	for current != nil {
		nCurrent.Random = oldToNew[current.Random]

		current = (*current).Next
		nCurrent = (*nCurrent).Next
	}

	return nHead
}

// func main() {
// 	values := []int{7, 13, 11, 10, 1}
// 	random := []int{-1, 0, 4, 2, 0}

// 	nodes := make([]Node, len(values))

// 	for ind, val := range values {
// 		nodes[ind] = Node{Val: val, Next: nil, Random: nil}
// 	}

// 	for i := 0; i < len(values)-1; i++ {
// 		nodes[i].Next = &nodes[i+1]
// 	}
// 	for i := 0; i < len(values); i++ {
// 		if random[i] != -1 {
// 			nodes[i].Random = &nodes[random[i]]
// 		}
// 	}

// 	_ = copyRandomList(&nodes[0])

// }
