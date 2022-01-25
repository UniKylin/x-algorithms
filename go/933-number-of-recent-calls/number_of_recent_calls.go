// link		: https://leetcode-cn.com/problems/number-of-recent-calls/
// Author	: Kylin
// Date		: 2022-01-25

package leetcode

type RecentCounter struct {
	queue []int
}

func Constructor() RecentCounter {
	return RecentCounter{
		queue: make([]int, 0),
	}
}

func (r *RecentCounter) Ping(t int) int {
	r.queue = append(r.queue, t)

	for r.queue[0] < t-3000 {
		r.queue = r.queue[1:len(r.queue)]
	}
	return len(r.queue)
}
