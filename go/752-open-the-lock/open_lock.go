// link		: https://leetcode-cn.com/problems/open-the-lock/
// Author	: Kylin
// Date		: 2022-01-18

package leetcode

import "strconv"

func OpenLock(deadends []string, target string) int {
	step := 0

	deadendsMap := make(map[string]bool)
	visitedMap := make(map[string]bool)

	for _, value := range deadends {
		deadendsMap[value] = true
	}

	q := []string{"0000"}

	for len(q) > 0 {
		size := len(q)

		for i := 0; i < size; i++ {
			node := q[0]
			q = q[1:]

			if node == target {
				return step
			}

			if _, visited := visitedMap[node]; visited {
				continue
			}

			if _, deadNode := deadendsMap[node]; deadNode {
				continue
			}

			visitedMap[node] = true

			for j := 0; j < len(node); j++ {
				num := int(node[j] - '0')
				up := (num + 1) % 10
				down := (num + 9) % 10
				q = append(q, node[:j]+strconv.Itoa(up)+node[j+1:])
				q = append(q, node[:j]+strconv.Itoa(down)+node[j+1:])
			}
		}
		step++
	}

	return -1
}
