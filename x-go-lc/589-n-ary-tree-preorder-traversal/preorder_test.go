// link		: https://leetcode.cn/problems/n-ary-tree-preorder-traversal
// Author	: Kylin
// Date		: 2022-12-26

package leetcode

import (
	"fmt"
	"testing"
)

func TestNPreorder(t *testing.T) {
	ant := Node{5, nil}
	bear := Node{6, nil}
	cat := Node{3, []*Node{&ant, &bear}}
	dog := Node{2, nil}
	egg := Node{4, nil}
	fox := Node{1, []*Node{&cat, &dog, &egg}}

	res := preorder(&fox)
	for index, data := range res {
		fmt.Println(index, data)
	}
}
