// link		: https://leetcode-cn.com/problems/same-tree/
// Author	: Kylin
// Date		: 2022-01-27

package leetcode

import (
	"fmt"
	"testing"
)

func TestSameTree(t *testing.T) {
	leftNode := &TreeNode{3, nil, nil}
	rightNode := &TreeNode{2, nil, nil}
	antRoot := &TreeNode{1, leftNode, rightNode}

	bearLeftNode := &TreeNode{3, nil, nil}
	bearRightNode := &TreeNode{2, nil, nil}
	bearRoot := &TreeNode{1, bearLeftNode, bearRightNode}

	res := SameTree(antRoot, bearRoot)
	fmt.Println(">>> is same: ", res)

}
