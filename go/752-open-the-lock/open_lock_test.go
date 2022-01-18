// link		: https://leetcode-cn.com/problems/open-the-lock/
// Author	: Kylin
// Date		: 2022-01-18

package leetcode

import (
	"fmt"
	"testing"
)

func TestOpenLock(t *testing.T) {
	deadends := []string{"0201", "0101", "0102", "1212", "2002"}
	target := "0202"

	res := OpenLock(deadends, target)
	fmt.Println(res)
}
