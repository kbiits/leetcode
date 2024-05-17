/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) int {
    return sumRootToLeafHelper(root, 0)
}

func sumRootToLeafHelper(root *TreeNode, sum int) int {
    if root == nil {
        return 0
    }

    sum = sum << 1 | root.Val

    // check leaf
    if root.Left == nil && root.Right == nil {
        return sum
    }

    return sumRootToLeafHelper(root.Left, sum) + sumRootToLeafHelper(root.Right, sum)
}
