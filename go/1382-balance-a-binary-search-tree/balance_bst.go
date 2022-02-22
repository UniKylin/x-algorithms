// link		: https://leetcode-cn.com/problems/balance-a-binary-search-tree/
// Author	: Kylin
// Date		: 2022-02-23

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// inorder + binary search
func BalanceBST(root *TreeNode) *TreeNode {
	res := make([]int, 0)

	var traversal func(node *TreeNode)
	traversal = func(node *TreeNode) {
		if node == nil {
			return
		}
		traversal(node.Left)
		res = append(res, node.Val)
		traversal(node.Right)
	}

	var build func(nums []int, left, right int) *TreeNode
	build = func(nums []int, left, right int) *TreeNode {
		if left > right {
			return nil
		}

		middle := left + (right-left)>>1
		root := &TreeNode{Val: nums[middle]}
		root.Left = build(nums, left, middle-1)
		root.Right = build(nums, middle+1, right)
		return root
	}

	traversal(root)
	return build(res, 0, len(res)-1)

}
