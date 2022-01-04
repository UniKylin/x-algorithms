// link		: https://leetcode-cn.com/problems/delete-node-in-a-bst/
// Author	: Kylin
// Date		: 2022-01-04

package leetcode

import (
	"fmt"
	"testing"
)

func TestDeleteNodeFromBST(t *testing.T) {
	root := &TreeNode{1, nil, nil}
	fmt.Println(root.Val)

	nextRoot := DeleteBST(root, 1)
	fmt.Println(nextRoot)
}
