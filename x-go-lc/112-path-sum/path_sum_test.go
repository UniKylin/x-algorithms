// link		: https://leetcode-cn.com/problems/path-sum/
// Author	: Kylin
// Date		: 2022-01-21

package leetcode

import (
	"fmt"
	"testing"
)

func TestPathSum(t *testing.T) {
	left := &TreeNode{Val: 2}
	right := &TreeNode{Val: 3}
	root := &TreeNode{Val: 1, Left: left, Right: right}

	targetSum := 3
	res := PathSum(root, targetSum)
	fmt.Println(">>> Has path:", res)
}
