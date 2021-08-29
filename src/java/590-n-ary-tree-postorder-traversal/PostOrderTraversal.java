// link		: https://leetcode-cn.com/problems/n-ary-tree-postorder-traversal/
// Author	: Kylin
// Date		: 2021-08-29

/*
// Definition for a Node.
class Node {
    public int val;
    public List<Node> children;

    public Node() {}

    public Node(int _val) {
        val = _val;
    }

    public Node(int _val, List<Node> _children) {
        val = _val;
        children = _children;
    }
};
*/

class PostOrderTraversal {
    List<Integer> result = new ArrayList<Integer>();

    public List<Integer> postorder(Node root) {
        if (root == null) {
            return result;
        }

        for (Node node : root.children) {
            postorder(node);
        }
        result.add(root.val);

        return result;
    }
}