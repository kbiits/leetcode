/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isUnivalTree(root *TreeNode) bool {
    return isUnivalTreeHelper(root, root.Val)
}

func isUnivalTreeHelper(root *TreeNode, val int) bool {
    if root == nil {
        return true
    }

    if root.Val != val {
        return false
    }

    return isUnivalTreeHelper(root.Left, val) && isUnivalTreeHelper(root.Right, val)
}
