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
    public boolean hasPathSum(TreeNode root, int targetSum) {
        return hasPathSum(root, 0, targetSum);
    }
    
    public boolean hasPathSum(TreeNode root, int curr, int targetSum) {
        if(root != null && root.left == null && root.right == null) {
            return curr + root.val == targetSum;
        }
        
        if(root == null) {
            return false;
        }
        
        return hasPathSum(root.left, curr + root.val, targetSum) || hasPathSum(root.right, curr + root.val, targetSum);
    }
}