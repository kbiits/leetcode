/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }

    balanced, _ := isBalancedHelper(root)
    return balanced
}

func isBalancedHelper(root *TreeNode) (bool, int) {
    if root == nil {
        return true, 0
    }

    isbalancedleft, left := isBalancedHelper(root.Left)
    isbalancedright, right := isBalancedHelper(root.Right)
    absDiff := left - right
    if absDiff < 0 {
        absDiff = -1 * absDiff
    }

    return isbalancedleft && isbalancedright && absDiff < 2 , max(left + 1, right + 1)
}

func max(a, b int) int {
    if a > b {
        return a
    }

    return b
}
