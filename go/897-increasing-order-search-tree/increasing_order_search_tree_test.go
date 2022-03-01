// link		: https://leetcode-cn.com/problems/increasing-order-search-tree/
// Author	: Kylin
// Date		: 2022-03-01

package leetcode

import (
	"fmt"
	"testing"
)

func TestIncreasingBST(t *testing.T) {
	right21 := &TreeNode{4, nil, nil}
	left21 := &TreeNode{2, nil, nil}
	right1 := &TreeNode{6, left21, right21}
	left1 := &TreeNode{3, left21, right21}
	root := &TreeNode{5, left1, right1}

	res := IncreasingBST(root)

	for cur := res; cur != nil; cur = cur.Right {
		fmt.Print(">>>", cur.Val)
	}
	fmt.Println()
}
