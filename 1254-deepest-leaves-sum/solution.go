/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func deepestLeavesSum(root *TreeNode) int {
    height := getHeight(root)
    return deepestLeavesSumHelper(root, height)
}

func deepestLeavesSumHelper(root *TreeNode, height int) int {
    if root == nil {
        return 0
    }

    if height == 1 {
        return root.Val
    }

    if height == 0 {
        return 0
    }

    left := deepestLeavesSumHelper(root.Left, height - 1)
    right := deepestLeavesSumHelper(root.Right, height - 1)
    return left + right
}

func getHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := getHeight(root.Left)
    right := getHeight(root.Right)
    if left > right {
        return 1 + left
    }

    return 1 + right
}
