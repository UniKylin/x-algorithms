// link		: https://leetcode-cn.com/problems/binary-tree-postorder-traversal/
// Author	: Kylin
// Date		: 2022-01-27

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PostorderTraversal(root *TreeNode) (result []int) {
	var postorder func(root *TreeNode)

	postorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		postorder(root.Left)
		postorder(root.Right)
		result = append(result, root.Val)
	}

	postorder(root)
	return
}
