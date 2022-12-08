// link		: https://leetcode-cn.com/problems/symmetric-tree/
// Author	: Kylin
// Date		: 2022-01-27

package leetcode

import (
	"fmt"
	"testing"
)

func TestSymmetric(t *testing.T) {
	leftNode := &TreeNode{3, nil, nil}
	rightNode := &TreeNode{2, nil, nil}
	root := &TreeNode{1, leftNode, rightNode}

	res := Symmetric(root)
	fmt.Println(res)

}
