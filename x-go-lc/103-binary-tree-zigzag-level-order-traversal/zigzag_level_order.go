// link		: https://leetcode.cn/problems/binary-tree-level-order-traversal
// Author	: Kylin
// Date		: 2022-12-24

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func zigzagLevelOrder(root *TreeNode) (arrs [][]int) {
	if root == nil {
		return
	}

	queue := []*TreeNode{root}
	for level := 0; len(queue) > 0; level++ {
		var items []int
		for i, length := 0, len(queue); i < length; i++ {
			curNode := queue[0]
			queue = queue[1:]

			items = append(items, curNode.Val)
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}

			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}

		if level%2 == 1 {
			items = reverse(items)
		}
		arrs = append(arrs, items)
	}

	return
}

func reverse(arr []int) []int {
	for left, right := 0, len(arr)-1; left < right; {
		arr[left], arr[right] = arr[right], arr[left]
		left++
		right--
	}
	return arr
}
