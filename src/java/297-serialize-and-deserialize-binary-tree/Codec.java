// link		: https://leetcode-cn.com/problems/serialize-and-deserialize-binary-tree/
// Author	: Kylin
// Date		: 2021-08-28

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode(int x) { val = x; }
 * }
 */
public class Codec {

    // Encodes a tree to a single string.
    public String serialize(TreeNode root) {
        return xserialize(root, "");
    }

    public String xserialize(TreeNode root, String str) {
        if (root == null) {
            return str += "None,";
        }

        str += str.valueOf(root.val) + ",";
        str = xserialize(root.left, str);
        str = xserialize(root.right, str);

        return str;
    }

    // Decodes your encoded data to tree.
    public TreeNode deserialize(String data) {
        String[] dataArr = data.split(",");
        List<String> dataList = new LinkedList<String>(Arrays.asList(dataArr));
        return xdeserialize(dataList);
    }

    public TreeNode xdeserialize(List<String> list) {
        if (list.get(0).equals("None")) {
            list.remove(0);
            return null;
        }

        TreeNode node = new TreeNode(Integer.valueOf(list.get(0)));
        list.remove(0);
        node.left = xdeserialize(list);
        node.right = xdeserialize(list);
        return node;
    }
}

// Your Codec object will be instantiated and called as such:
// Codec ser = new Codec();
// Codec deser = new Codec();
// TreeNode ans = deser.deserialize(ser.serialize(root));