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
    public List<TreeNode> allPossibleFBT(int n) {
        List<TreeNode> list = new ArrayList<>();
        
        if(n <= 0 || n % 2 == 0) {
            return list;
        }
        
        if(n == 1) {
            list.add(new TreeNode());
            return list;
        }
        
        for(int i = 1; i < n; i += 2) {
            int rightNum = n - i - 1;
            
            List<TreeNode> leftSubTree = allPossibleFBT(i);
            List<TreeNode> rightSubTree = allPossibleFBT(rightNum);
            
            for(TreeNode l : leftSubTree) {
                for(TreeNode r: rightSubTree) {
                    TreeNode root = new TreeNode();
                    root.left = l;
                    root.right = r;
                    list.add(root);
                }
            }
            
        }
        
        return list;
    }
}
