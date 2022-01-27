// link		: https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
// Author	: Kylin
// Date		: 2022-01-27

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PreorderTraversal(root *TreeNode) (result []int) {
	var preorder func(root *TreeNode)

	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		result = append(result, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)
	return
}
