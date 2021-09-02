// link		: https://leetcode-cn.com/problems/minimum-path-sum/
// Author	: Kylin
// Date		: 2021-09-02

package leetcode

func minPathSum(grid [][]int) int {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			if row == 0 && column == 0 {
				continue
			} else if row == 0 {
				grid[row][column] = grid[row][column-1] + grid[row][column]
			} else if column == 0 {
				grid[row][column] = grid[row-1][column] + grid[row][column]
			} else {
				low := 0
				if grid[row-1][column] <= grid[row][column-1] {
					low = grid[row-1][column]
				} else {
					low = grid[row][column-1]
				}
				grid[row][column] = low + grid[row][column]
			}
		}
	}

	return grid[len(grid)-1][len(grid[0])-1]
}
