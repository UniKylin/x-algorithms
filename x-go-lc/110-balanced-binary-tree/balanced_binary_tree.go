// link		: https://leetcode-cn.com/problems/balanced-binary-tree/
// Author	: Kylin
// Date		: 2022-02-23

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func IsBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}

	return abs(depth(root.Left)-depth(root.Right)) <= 1 && IsBalanced(root.Left) && IsBalanced(root.Right)
}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}

	return max(depth(root.Left), depth(root.Right)) + 1
}

func max(ant, bear int) int {
	if ant > bear {
		return ant
	}
	return bear
}

func abs(ant int) int {
	if ant < 0 {
		return -ant
	}
	return ant
}
