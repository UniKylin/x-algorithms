// link		: https://leetcode-cn.com/problems/unique-paths-ii/
// Author	: Kylin
// Date		: 2022-01-06

package leetcode

import (
	"fmt"
	"testing"
)

func TestUniquePathsWithObstacles(t *testing.T) {
	obstacleGrid := [][]int{
		{0, 0, 0},
		{0, 1, 0},
		{0, 0, 0},
	}

	res := UniquePathsWithObstacles(obstacleGrid)
	fmt.Println(">>> res:", res)
}
