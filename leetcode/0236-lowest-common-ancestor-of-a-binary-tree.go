package main

func dfs(root, p, q *TreeNode) (pFound, qFound bool, ans *TreeNode) {
	if root == nil {
		return false, false, nil
	}
	if root.Val == p.Val {
		pFound = true
	}

	if root.Val == q.Val {
		qFound = true
	}

	if pFound && qFound {
		return true, true, root
	}
	var _pFound, _qFound, _ans = dfs(root.Left, p, q)
	if _pFound && _qFound {
		return true, true, _ans
	}
	pFound = pFound || _pFound
	qFound = qFound || _qFound
	if pFound && qFound {
		return true, true, root
	}
	_pFound, _qFound, _ans = dfs(root.Right, p, q)
	if _pFound && _qFound {
		return true, true, _ans
	}
	pFound = pFound || _pFound
	qFound = qFound || _qFound
	if pFound && qFound {
		return true, true, root
	}

	return
}
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var _, _, ans = dfs(root, p, q)
	if ans != nil {
		return ans
	}

	return root
}

//
//func main() {
//	var a *TreeNode = &TreeNode{Val: 7}
//	var b *TreeNode = &TreeNode{Val: 4}
//	var c *TreeNode = &TreeNode{Val: 2, Left: a, Right: b}
//	var d *TreeNode = &TreeNode{Val: 6}
//	var e *TreeNode = &TreeNode{Val: 5, Left: d, Right: c}
//	var f *TreeNode = &TreeNode{Val: 0}
//	var g *TreeNode = &TreeNode{Val: 8}
//	var h *TreeNode = &TreeNode{Val: 1, Left: f, Right: g}
//	var i *TreeNode = &TreeNode{Val: 3, Left: e, Right: h}
//
//	var res = lowestCommonAncestor(i, e, g)
//	fmt.Println(res.Val)
//}
