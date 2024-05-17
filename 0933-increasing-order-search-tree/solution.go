/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func increasingBST(root *TreeNode) *TreeNode {

    var prevTraversed *TreeNode
    var recursive func(root *TreeNode)
    recursive = func(root *TreeNode) {
        if root.Right != nil {
            recursive(root.Right)
        }
        root.Right = prevTraversed
        prevTraversed = root
        if root.Left != nil {
            recursive(root.Left)
            root.Left = nil
        }
    }
    recursive(root)

    return prevTraversed
}
