/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func diameterOfBinaryTree(root *TreeNode) int {
    if root == nil {
        return 0
    }

    max := 0
    var dfs func(root *TreeNode) int
    dfs = func (root *TreeNode) int {
        if root == nil {
            return 0
        }

        left := dfs(root.Left)
        right := dfs(root.Right)

        if left + right > max {
            max = left + right
        }

        return 1 + maxFunc(left, right)
    }

    dfs(root)
    return max
}

func maxFunc(a,b int) int {
    if a > b {
        return a
    }
    
    return b
}
