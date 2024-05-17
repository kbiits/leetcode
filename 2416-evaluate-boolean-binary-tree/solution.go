/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func evaluateTree(root *TreeNode) bool {
    if root.Val == 1 || root.Val == 0 {
        return root.Val == 1
    }

    if root.Val == 2 {
        return evaluateTree(root.Left) || evaluateTree(root.Right)
    }

    return evaluateTree(root.Left) && evaluateTree(root.Right)
}
