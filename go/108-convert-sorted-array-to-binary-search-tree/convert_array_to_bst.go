// link		: https://leetcode-cn.com/problems/convert-sorted-array-to-binary-search-tree/
// Author	: Kylin
// Date		: 2022-02-23

package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func SortedArrayToBST(nums []int) *TreeNode {
	return build(nums, 0, len(nums)-1)
}

func build(nums []int, left, right int) *TreeNode {
	if left > right {
		return nil
	}

	middle := left + (right-left+1)>>1
	root := &TreeNode{Val: nums[middle]}
	root.Left = build(nums, left, middle-1)
	root.Right = build(nums, middle+1, right)

	return root
}
