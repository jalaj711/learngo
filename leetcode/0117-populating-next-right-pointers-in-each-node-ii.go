package main

type _Node struct {
	Val   int
	Left  *_Node
	Right *_Node
	Next  *_Node
}

func connect(root *_Node) *_Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}

	q := make([]*_Node, 1)
	q[0] = root
	n := 0
	orig := 0
	for len(q) != 0 {
		orig = len(q)
		for n = 0; n < orig; n++ {
			if q[n].Left != nil {
				q = append(q, q[n].Left)
			}
			if q[n].Right != nil {
				q = append(q, q[n].Right)
			}
			if n < orig-1 {
				q[n].Next = q[n+1]
			}
		}
		q = q[orig:]
	}

	return root
}

// func main() {
// 	ll := _Node{Val: 4}
// 	lr := _Node{Val: 5}
// 	rr := _Node{Val: 7}
// 	l := _Node{Val: 2, Left: &ll, Right: &lr}
// 	r := _Node{Val: 3, Right: &rr}
// 	root := _Node{Val: 1, Left: &l, Right: &r}

// 	connect(&root)
// }
