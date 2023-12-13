// link		: https://leetcode-cn.com/problems/path-sum/
// Author	: Kylin
// Date		: 2022-01-21

package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}

	queue := []*TreeNode{root}

	for len(queue) != 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Left == nil && node.Right == nil && node.Val == targetSum {
			return true
		}

		if node.Left != nil {
			node.Left.Val += node.Val
			queue = append(queue, node.Left)
		}

		if node.Right != nil {
			node.Right.Val += node.Val
			queue = append(queue, node.Right)
		}
	}

	return false
}
