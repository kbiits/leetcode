/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }

    leftMax := maxDepth(root.Left)
    rightMax := maxDepth(root.Right)

    if leftMax > rightMax {
        return 1 + leftMax
    }

    return 1 + rightMax
}
