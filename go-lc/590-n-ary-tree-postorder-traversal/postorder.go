// link		: https://leetcode.cn/problems/n-ary-tree-postorder-traversal/
// Author	: Kylin
// Date		: 2023-03-14

package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) (res []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}

		for _, childNode := range node.Children {
			dfs(childNode)
		}
		res = append(res, node.Val)
	}

	dfs(root)
	return
}
