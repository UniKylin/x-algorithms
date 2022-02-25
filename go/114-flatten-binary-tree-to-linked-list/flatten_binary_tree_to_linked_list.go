// link		: https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
// Author	: Kylin
// Date		: 2022-02-25

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Flatten(root *TreeNode) {
	list := preorderTraversal(root)

	for i := 1; i < len(list); i++ {
		prev, cur := list[i-1], list[i]
		prev.Left, prev.Right = nil, cur
	}
}

func preorderTraversal(root *TreeNode) []*TreeNode {
	list := []*TreeNode{}

	if root != nil {
		list = append(list, root)
		list = append(list, preorderTraversal(root.Left)...)
		list = append(list, preorderTraversal(root.Right)...)
	}

	return list
}
