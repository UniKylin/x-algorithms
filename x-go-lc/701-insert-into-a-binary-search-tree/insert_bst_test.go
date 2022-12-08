// link		: https://leetcode-cn.com/problems/insert-into-a-binary-search-tree/
// Author	: Kylin
// Date		: 2022-01-04

package leetcode

import (
	"fmt"
	"testing"
)

func TestInsertBST(t *testing.T) {
	value := 7
	root := TreeNode{4, nil, nil}
	newRoot := InsertIntoBST(&root, value)

	// fmt.Println(newRoot.Val)
	// fmt.Println(newRoot.Right.Val)

	nextValue := 2
	nextRoot := InsertIntoBST(newRoot, nextValue)
	fmt.Println(nextRoot.Val)
	fmt.Println(nextRoot.Left.Val)
	fmt.Println(nextRoot.Right.Val)
}
