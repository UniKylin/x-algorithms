// link		: https://leetcode-cn.com/problems/find-largest-value-in-each-tree-row/
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
    public List<Integer> largestValues(TreeNode root) {
        List<Integer> result = new ArrayList();
        Queue<TreeNode> queue = new ArrayDeque();

        if (root != null) {
            queue.add(root);
        }

        while (!queue.isEmpty()) {
            // So Integer.MIN_VALUE is the bigest number of Integer
            int max = Integer.MIN_VALUE;
            int count = queue.size();
            for (int i = 0; i < count; i++) {
                TreeNode node = queue.poll();
                max = node.val > max ? node.val : max;

                if (node.left != null) {
                    queue.add(node.left);
                }

                if (node.right != null) {
                    queue.add(node.right);
                }
            }
            result.add(max);
        }

        return result;
    }
}