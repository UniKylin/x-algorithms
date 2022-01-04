// link		: https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
// Author	: Kylin
// Date		: 2022-01-04

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InsertIntoBST(root *TreeNode, val int) *TreeNode {
	// 1. 如果是一个空树那就创建节点
	if root == nil {
		return &TreeNode{Val: val}
	}

	// 2. 不为空，根据二叉搜索树性质判断
	cur := root

	for cur != nil {
		if val < cur.Val {
			if cur.Left == nil {
				cur.Left = &TreeNode{Val: val}
				break
			}
			cur = cur.Left
		} else {
			if cur.Right == nil {
				cur.Right = &TreeNode{Val: val}
				break
			}
			cur = cur.Right
		}
	}

	return root
}
