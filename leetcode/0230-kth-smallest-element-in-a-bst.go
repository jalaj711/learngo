package main

func _processElem(root *TreeNode, k *int) (ans int) {
	ans = -1
	if root == nil {
		return ans
	}
	if root.Left != nil {
		ans = _processElem(root.Left, k)
	}
	if *k == 0 && ans != -1 {
		return ans
	}
	*k--
	if *k == 0 {
		return root.Val
	}
	if root.Right != nil {
		return _processElem(root.Right, k)
	}
	return -1
}

func kthSmallest(root *TreeNode, k int) int {
	return _processElem(root, &k)
}
