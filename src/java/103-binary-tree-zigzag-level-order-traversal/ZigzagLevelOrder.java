// link		: https://leetcode-cn.com/problems/binary-tree-zigzag-level-order-traversal/
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
    public List<List<Integer>> zigzagLevelOrder(TreeNode root) {
        List<List<Integer>> result = new ArrayList();
        Queue<TreeNode> queue = new ArrayDeque();

        if (root != null) {
            queue.add(root);
        }

        // 控制入队方向
        boolean left2Right = true;

        while (!queue.isEmpty()) {
            List<Integer> level = new ArrayList();
            int count = queue.size();

            for (int i = 0; i < count; i++) {
                TreeNode node = queue.poll();

				if (left2Right) {
                    level.add(node.val);
                } else {
                    level.add(0, node.val);
                }

                if (node.left != null) {
                    queue.add(node.left);
                }

                if (node.right != null) {
                    queue.add(node.right);
                }
            }
            left2Right = !left2Right;
            result.add(level);
        }

        return result;
    }
}