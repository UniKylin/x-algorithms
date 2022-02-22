// link		: https://leetcode-cn.com/problems/balanced-binary-tree/
// Author	: Kylin
// Date		: 2022-02-23

package leetcode

import (
	"fmt"
	"testing"
)

func TestBalancedBinaryTree(t *testing.T) {
	right31 := &TreeNode{5, nil, nil}
	left31 := &TreeNode{3, nil, nil}

	right21 := &TreeNode{4, left31, right31}
	left21 := &TreeNode{0, nil, nil}

	left1 := &TreeNode{2, left21, right21}
	right1 := &TreeNode{8, nil, nil}
	root := &TreeNode{6, left1, right1}

	res := IsBalanced(root)

	fmt.Println(">>> Balanced:", res)

}
