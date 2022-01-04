// link		: https://leetcode-cn.com/problems/search-in-a-binary-search-tree/
// Author	: Kylin
// Date		: 2022-01-04

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SearchBST(root *TreeNode, val int) *TreeNode {
	for root != nil {
		if val == root.Val {
			return root
		}

		if val < root.Val {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return nil
}
