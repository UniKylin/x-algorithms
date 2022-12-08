// link		: https://leetcode-cn.com/problems/design-browser-history/
// Author	: Kylin
// Date		: 2021-12-27

package leetcode

type BrowserHistory struct {
	cur   []string
	index int
}

func BrowserHistoryConstructor(homepage string) BrowserHistory {
	return BrowserHistory{
		index: 0,
		cur:   []string{homepage},
	}
}

func (b *BrowserHistory) Visit(url string) {
	count := len(b.cur)

	if b.index == count-1 {
		b.cur = append(b.cur, url)
	} else if b.index < count-1 {
		b.cur = b.cur[:b.index+1]
		b.cur = append(b.cur, url)
	}
	b.index++
}

func (b *BrowserHistory) Back(steps int) string {
	if steps > b.index {
		b.index = 0
		return b.cur[0]
	}
	b.index -= steps
	return b.cur[b.index]
}

func (b *BrowserHistory) Forward(steps int) string {
	count := len(b.cur)

	if b.index == count-1 {
		return b.cur[b.index]
	} else if b.index+steps > count-1 {
		b.index = count - 1
		return b.cur[count-1]
	} else {
		b.index += steps
		return b.cur[b.index]
	}
}

/**
 * Your BrowserHistory object will be instantiated and called as such:
 * obj := BrowserHistoryConstructor(homepage);
 * obj.Visit(url);
 * param_2 := obj.Back(steps);
 * param_3 := obj.Forward(steps);
 */
