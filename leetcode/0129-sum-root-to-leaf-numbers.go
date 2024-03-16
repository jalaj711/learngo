package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func visitNode(node *TreeNode, curr *int, ans *int) {
	*curr = *curr*10 + node.Val
	if node.Left == nil && node.Right == nil {
		*ans += *curr
		// fmt.Printf("found path: %d, ans=%d", *curr, *ans)
	} else {
		if node.Left != nil {
			visitNode(node.Left, curr, ans)
		}
		if node.Right != nil {
			visitNode(node.Right, curr, ans)
		}
	}
	*curr = (*curr) / 10
}

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	curr := 0
	visitNode(root, &curr, &ans)
	return ans
}
