// link		: https://leetcode-cn.com/problems/binary-tree-right-side-view/
// Author	: Kylin
// Date		: 2021-08-28

/**
 * Definition for a binary tree node.
 * public class TreeNode {
 *     int val;
 *     TreeNode left;
 *     TreeNode right;
 *     TreeNode() {}
 *     TreeNode(int val) { this.val = val; }
 *     TreeNode(int val, TreeNode left, TreeNode right) {
 *         this.val = val;
 *         this.left = left;
 *         this.right = right;
 *     }
 * }
 */
class Solution {
    public List<Integer> rightSideView(TreeNode root) {
        List<Integer> result = new ArrayList();
        Queue<TreeNode> queue = new ArrayDeque();

        if (root != null) {
            queue.add(root);
        }

        while (!queue.isEmpty()) {
            int count = queue.size();

            for (int i = 0; i < count; i++) {
                TreeNode node = queue.poll();

                if (node.left != null) {
                    queue.add(node.left);
                }

                if (node.right != null) {
                    queue.add(node.right);
                }

                if (i == count -1) {
                    result.add(node.val);
                }
            }
        }

        return result;
    }
}