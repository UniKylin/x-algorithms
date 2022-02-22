// link		: https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree-iii/
// Author	: Kylin
// Date		: 2022-02-23

package leetcode

type Node struct {
	Val    int
	Left   *Node
	Right  *Node
	Parent *Node
}

func LowestCommonAncestor(p *Node, q *Node) *Node {
	curP := p
	curQ := q

	for curP != curQ {
		if curP == nil {
			curP = q
		} else {
			curP = curP.Parent
		}

		if curQ == nil {
			curQ = p
		} else {
			curQ = curQ.Parent
		}
	}

	return curP
}
