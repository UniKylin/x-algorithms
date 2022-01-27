// link		: https://leetcode-cn.com/problems/binary-tree-preorder-traversal/
// Author	: Kylin
// Date		: 2022-01-27

package leetcode

import (
	"fmt"
	"testing"
)

func TestInorderTraversal(t *testing.T) {
	leftNode := &TreeNode{3, nil, nil}
	rightNode := &TreeNode{2, leftNode, nil}
	root := &TreeNode{1, nil, rightNode}

	res := PreorderTraversal(root)
	fmt.Println(res)

	onlyRoot := &TreeNode{1, nil, nil}
	onlyRes := PreorderTraversal(onlyRoot)
	fmt.Println(onlyRes)
}
