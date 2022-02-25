// link		: https://leetcode-cn.com/problems/flatten-binary-tree-to-linked-list/
// Author	: Kylin
// Date		: 2022-02-25

package leetcode

import (
	"testing"
)

func TestFlatten(t *testing.T) {
	right22 := &TreeNode{6, nil, nil}

	right21 := &TreeNode{4, nil, nil}
	left21 := &TreeNode{3, nil, nil}

	right1 := &TreeNode{5, nil, right22}
	left1 := &TreeNode{2, left21, right21}
	root := &TreeNode{1, left1, right1}

	Flatten(root)
}
