// link		: https://leetcode-cn.com/problems/increasing-order-search-tree/
// Author	: Kylin
// Date		: 2022-03-01

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IncreasingBST(root *TreeNode) *TreeNode {
	res := make([]int, 0)

	var inorder func(*TreeNode)
	inorder = func(root *TreeNode) {
		if root == nil {
			return
		}
		inorder(root.Left)
		res = append(res, root.Val)
		inorder(root.Right)
	}

	inorder(root)

	dummy := &TreeNode{}
	cur := dummy

	for _, val := range res {
		cur.Right = &TreeNode{Val: val}
		cur = cur.Right
	}

	return dummy.Right
}
