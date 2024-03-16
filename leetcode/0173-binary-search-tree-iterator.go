package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type __Node struct {
	Val  *TreeNode
	Prev *__Node
}

type BSTIterator struct {
	stack *__Node
}

func ConstructorBST(root *TreeNode) BSTIterator {
	iterator := BSTIterator{stack: &__Node{Val: root, Prev: nil}}
	iterator.leftChildTraversal(root.Left)
	return iterator
}

func (this *BSTIterator) leftChildTraversal(root *TreeNode) {
	if root == nil {
		return
	}

	this.stack = &__Node{Val: root, Prev: this.stack}
	this.leftChildTraversal(root.Left)
}

func (this *BSTIterator) Next() int {
	ans := this.stack.Val
	this.stack = this.stack.Prev
	this.leftChildTraversal(ans.Right)
	return ans.Val
}

func (this *BSTIterator) HasNext() bool {
	return this.stack != nil
}

/**
 * Your BSTIterator object will be instantiated and called as such:
 * obj := Constructor(root);
 * param_1 := obj.Next();
 * param_2 := obj.HasNext();
 */
