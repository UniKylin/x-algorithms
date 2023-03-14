// link		: https://leetcode.cn/problems/n-ary-tree-level-order-traversal/
// Author	: Kylin
// Date		: 2023-03-15

package leetcode

import (
	"fmt"
	"testing"
)

func TestLevelOrder(t *testing.T) {

	ant := Node{5, nil}
	bear := Node{6, nil}
	cat := Node{3, []*Node{&ant, &bear}}
	dog := Node{2, nil}
	egg := Node{4, nil}
	fox := Node{1, []*Node{&cat, &dog, &egg}}

	res := levelOrder(&fox)
	for index, data := range res {
		fmt.Println(index, data)
	}
}
