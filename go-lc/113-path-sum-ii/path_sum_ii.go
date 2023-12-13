// link		: https://leetcode-cn.com/problems/path-sum-ii/
// Author	: Kylin
// Date		: 2022-01-23

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func PathSumII(root *TreeNode, targetSum int) [][]int {
	res := [][]int{}

	var dfs func(root *TreeNode, sum int, stack []int)
	dfs = func(root *TreeNode, sum int, stack []int) {
		if root == nil {
			return
		}

		stack = append(stack, root.Val)

		if root.Left == nil && root.Right == nil {
			if sum == root.Val {
				temp := make([]int, len(stack))
				copy(temp, stack)
				res = append(res, temp)
			}
		}

		dfs(root.Left, sum-root.Val, stack)
		dfs(root.Right, sum-root.Val, stack)
	}

	dfs(root, targetSum, []int{})

	return res
}
