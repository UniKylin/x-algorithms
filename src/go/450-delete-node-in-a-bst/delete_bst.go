// link		: https://leetcode-cn.com/problems/delete-node-in-a-bst/
// Author	: Kylin
// Date		: 2022-01-04

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 二叉搜索树删除情况比较多分类讨论
func DeleteBST(root *TreeNode, val int) *TreeNode {
	// 1. 空节点直接返回
	if root == nil {
		return nil
	}

	// 2. 查找到节点的情况
	if root.Val == val {
		// 2.1 左子树右子树都为空 直接删除节点
		if root.Left == nil && root.Right == nil {
			return nil
		}

		// 2.2 左空右不空，删除节点，右子树补上去
		if root.Left == nil && root.Right != nil {
			root = root.Right
			return root
		}

		// 2.3 左不空右空
		if root.Left != nil && root.Right == nil {
			root = root.Left
			return root
		}

		// 2.3 左右子树都不为空，则将删除节点的左子树放到被删除节点右子树上最左边的左子树上的位置
		// 临时保存左右子树
		leftTree := root.Left
		rightTree := root.Right

		// 遍历找到右子树上最左下角的左子树位置
		temp := root.Right

		for temp.Left != nil {
			temp = temp.Left
		}

		// 讲左子树挂在到最左下角的子树上
		temp.Left = leftTree
		// 更新root节点，相当于删除root节点
		root = rightTree

		return root
	}

	// 3.查找不到节点的情况，递归左右子树继续找
	if val < root.Val {
		root.Left = DeleteBST(root.Left, val)
	}

	if val > root.Val {
		root.Right = DeleteBST(root.Right, val)
	}

	return root
}
