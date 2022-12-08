// link		: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// Author	: Kylin
// Date		: 2022-02-22

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func LowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	cur := root

	for {
		if p.Val < cur.Val && q.Val < cur.Val {
			cur = cur.Left
		} else if p.Val > cur.Val && q.Val > cur.Val {
			cur = cur.Right
		} else {
			return cur
		}
	}

}
