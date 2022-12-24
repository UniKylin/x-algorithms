// link		: https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
// Author	: Kylin
// Date		: 2022-12-24

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) (res [][]int) {
	// queue
	queue := []*TreeNode{}
	if root != nil {
		queue = append(queue, root)
	}

	for len(queue) > 0 {
		// level node count
		count := len(queue)

		// current level nodes
		var level []int

		for i := 0; i < count; i++ {
			node := queue[0]
			queue = queue[1:]

			level = append(level, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, level)
	}

	return
}
