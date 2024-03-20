package main

func zigzagLevelOrder(root *TreeNode) [][]int {
	ans := make([][]int, 0, 1)
	if root == nil {
		return ans
	}

	q := make([]*TreeNode, 1)
	q[0] = root
	n := 0
	left := true
	for len(q) != 0 {
		n = len(q)
		row := make([]int, n)
		for i := 0; i < n; i++ {
			if left {
				row[i] = q[i].Val
			} else {
				row[n-i-1] = q[i].Val
			}
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		left = !left
		ans = append(ans, row)
		q = q[n:]
	}
	return ans
}
