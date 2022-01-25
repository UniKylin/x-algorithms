// link		: https://leetcode-cn.com/problems/number-of-recent-calls/
// Author	: Kylin
// Date		: 2022-01-25

package leetcode

import (
	"fmt"
	"testing"
)

func TestRecentCalls(t *testing.T) {
	r := Constructor()

	ant := r.Ping(1)
	fmt.Println(">>>", ant)

	bear := r.Ping(1000)
	fmt.Println(">>>", bear)

	cat := r.Ping(3001)
	fmt.Println(">>>", cat)

	dog := r.Ping(3003)
	fmt.Println(">>>", dog)

}
