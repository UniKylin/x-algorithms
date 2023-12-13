// link		: https://leetcode-cn.com/problems/symmetric-tree/
// Author	: Kylin
// Date		: 2022-01-27

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Symmetric(root *TreeNode) bool {
	return process(root, root)
}

func process(p, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	}

	if p == nil || q == nil {
		return false
	}

	return p.Val == q.Val && process(p.Left, q.Right) && process(p.Right, q.Left)
}
