// link		: https://leetcode.cn/problems/n-ary-tree-level-order-traversal/
// Author	: Kylin
// Date		: 2023-03-15

package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) (res [][]int) {
	if root == nil {
		return
	}

	queue := []*Node{}
	queue = append(queue, root)

	for len(queue) > 0 {
		level := []int{}
		temp := queue
		queue = nil

		for _, node := range temp {
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}
		res = append(res, level)
	}

	return
}
