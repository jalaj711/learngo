package main

func _flatten(root *TreeNode) (*TreeNode, *TreeNode) {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root, root
	}
	lhead, ltail := _flatten(root.Left)
	rhead, rtail := _flatten(root.Right)

	root.Left = nil
	if ltail != nil {
		ltail.Right = rhead
		root.Right = lhead
		if rtail == nil {
			return root, ltail
		}
	} else {
		root.Right = rhead
	}
	return root, rtail
}

func flatten(root *TreeNode) {
	_flatten(root)
}

// func main() {
// 	// ll := TreeNode{Val: 4}
// 	// lr := TreeNode{Val: 5}
// 	// rr := TreeNode{Val: 7}
// 	// l := TreeNode{Val: 2, Left: &ll, Right: &lr}
// 	// r := TreeNode{Val: 3, Right: &rr}
// 	// root := TreeNode{Val: 1, Left: &l, Right: &r}

// 	ll := TreeNode{Val: 3}
// 	l := TreeNode{Val: 2, Left: &ll}
// 	root := TreeNode{Val: 1, Left: &l}
// 	flatten(&root)
// }
