// link		: https://leetcode-cn.com/problems/house-robber/
// Author	: Kylin
// Date		: 2021-12-29

package leetcode

import (
	"fmt"
	"testing"
)

func TestRobber(t *testing.T) {
	firstHouse := []int{1, 2, 3, 1} //4
	firstRes := Robber(firstHouse)
	fmt.Println(firstRes)

	anotherHouse := []int{2, 7, 9, 3, 1} //12
	anotherRes := Robber(anotherHouse)
	fmt.Println(anotherRes)

}
