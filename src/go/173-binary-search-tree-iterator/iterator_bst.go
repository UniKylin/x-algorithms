// link		: https://leetcode-cn.com/problems/binary-search-tree-iterator/
// Author	: Kylin
// Date		: 2022-01-04

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type BSTIterator struct {
	arr []int
}

func Constructor(root *TreeNode) (bst BSTIterator) {
	bst.inorder(root)
	return
}

func (bst *BSTIterator) inorder(node *TreeNode) {
	if node == nil {
		return
	}

	bst.inorder(node.Left)
	bst.arr = append(bst.arr, node.Val)
	bst.inorder(node.Right)
}

func (bst *BSTIterator) Next() int {
	val := bst.arr[0]
	bst.arr = bst.arr[1:]
	return val
}

func (bst *BSTIterator) HasNext() bool {
	return len(bst.arr) > 0
}
