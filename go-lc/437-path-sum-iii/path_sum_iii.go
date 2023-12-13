// link		: https://leetcode-cn.com/problems/path-sum-iii/
// Author	: Kylin
// Date		: 2022-01-23

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSumIII(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}

	res := rootSum(root, targetSum)
	res += PathSumIII(root.Left, targetSum)
	res += PathSumIII(root.Right, targetSum)

	return res
}

func rootSum(root *TreeNode, targetSum int) (result int) {
	if root == nil {
		return
	}

	if root.Val == targetSum {
		result++
	}

	result += rootSum(root.Left, targetSum-root.Val)
	result += rootSum(root.Right, targetSum-root.Val)

	return
}
