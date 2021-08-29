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

class PreOrderTraversal {

    List<Integer> result = new ArrayList<Integer>();

    public List<Integer> preorder(Node root) {

        if (root == null) {
            return result;
        }

        result.add(root.val);

        for (Node node : root.children) {
            preorder(node);
        }

        return result;
    }
}