// link		: https://leetcode-cn.com/problems/binary-tree-inorder-traversal/
// Author	: Kylin
// Date		: 2022-01-27

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func InorderTraversal(root *TreeNode) (result []int) {
	var inorder func(root *TreeNode)

	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		result = append(result, root.Val)
		inorder(root.Right)
	}

	inorder(root)
	return
}
