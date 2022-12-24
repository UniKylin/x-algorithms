package leetcode

import (
	"fmt"
	"testing"
)

func TestLevelOrderTraversal(t *testing.T) {
	leftNode := &TreeNode{3, nil, nil}
	rightNode := &TreeNode{2, leftNode, nil}
	root := &TreeNode{1, nil, rightNode}

	res := levelOrder(root)
	fmt.Println(res)
}
