/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func goodNodes(root *TreeNode) int {
    return goodNodesHelper(root, -1e5)
}

func goodNodesHelper(root *TreeNode, maxVal int) int {
    if root == nil {
        return 0
    }

    if maxVal > root.Val {
        return goodNodesHelper(root.Left, maxVal) + goodNodesHelper(root.Right, maxVal)
    }

    return 1 + goodNodesHelper(root.Left, root.Val) + goodNodesHelper(root.Right, root.Val)
}
