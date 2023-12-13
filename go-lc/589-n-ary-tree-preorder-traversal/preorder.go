// link		: https://leetcode.cn/problems/n-ary-tree-preorder-traversal
// Author	: Kylin
// Date		: 2022-12-26

package leetcode

type Node struct {
	Val      int
	Children []*Node
}

func preorder(root *Node) (res []int) {
	var dfs func(node *Node)
	dfs = func(node *Node) {
		if node == nil {
			return
		}

		res = append(res, node.Val)
		for _, childNode := range node.Children {
			dfs(childNode)
		}
	}

	dfs(root)
	return
}
