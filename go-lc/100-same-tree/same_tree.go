// link		: https://leetcode-cn.com/problems/same-tree/
// Author	: Kylin
// Date		: 2022-01-27

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	if p.Val != q.Val {
		return false
	}

	return SameTree(p.Left, q.Left) && SameTree(p.Right, q.Right)
}
