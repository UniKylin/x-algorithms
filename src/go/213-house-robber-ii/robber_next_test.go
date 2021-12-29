// link		: https://leetcode-cn.com/problems/house-robber-ii/
// Author	: Kylin
// Date		: 2021-12-29

package leetcode

import (
	"fmt"
	"testing"
)

func TestRobberNext(t *testing.T) {
	firstHouse := []int{1, 2, 3, 1} //4
	firstRes := RobberNext(firstHouse)
	fmt.Println(firstRes)

	anotherHouse := []int{2, 7, 9, 3, 1} //11
	anotherRes := RobberNext(anotherHouse)
	fmt.Println(anotherRes)
}
