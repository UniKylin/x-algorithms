// link		: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-search-tree/
// Author	: Kylin
// Date		: 2022-02-22

package leetcode

import (
	"fmt"
	"testing"
)

func TestLowestCommonAncestor(t *testing.T) {
	right31 := &TreeNode{5, nil, nil}
	left31 := &TreeNode{3, nil, nil}

	right22 := &TreeNode{9, nil, nil}
	right21 := &TreeNode{4, left31, right31}
	left22 := &TreeNode{7, nil, nil}
	left21 := &TreeNode{0, nil, nil}

	left1 := &TreeNode{2, left21, right21}
	right1 := &TreeNode{8, left22, right22}
	root := &TreeNode{6, left1, right1}

	p := &TreeNode{2, nil, nil}
	q := &TreeNode{8, nil, nil}

	node := LowestCommonAncestor(root, p, q)
	fmt.Println(">>> Result node:", node.Val)
}
