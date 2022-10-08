/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {string[]}
 */
var binaryTreePaths = function(root) {
    const arr = [];
    
    dfs(root, "", arr);
    
    return arr;
};

const dfs = (root, str = "", array = []) => {
    if (root == null) {
        return;
    }
    
    str += root.val == 0 ? "" : root.val.toString() + "->";
    dfs(root.left, str, array);
    dfs(root.right, str, array);
    
    
    if (!root.left && !root.right) {
        array.push(str.substring(0, str.length - 2));    
    }
    
    return;
}