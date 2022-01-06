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
    public TreeNode searchBST(TreeNode root, int val) {
        if(root == null) {
            return null;
        }
        
        if(root.val == val) {
            return root;
        }
        
        
        TreeNode res = searchBST(root.left, val);
        if(res != null) {
            return res;
        }
        TreeNode resR = searchBST(root.right, val);
        if(resR != null) {
            return resR;
        }
        
        return null;
    }
}