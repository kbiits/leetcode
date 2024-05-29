/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isValidBST(root *TreeNode) bool {
    if root == nil {
        return true
    }

    if root.Left == nil && root.Right == nil {
        return true
    }

    if root.Left != nil {
        if getRightestNodeVal(root.Left) >= root.Val {
            return false
        }
    }

    if root.Right != nil {
        if getLeftestNodeVal(root.Right) <= root.Val {
            return false
        }
    }

    return isValidBST(root.Left) && isValidBST(root.Right)
}

func getRightestNodeVal(root *TreeNode) int {
    if root.Right == nil {
        return root.Val
    }

    return getRightestNodeVal(root.Right)
}

func getLeftestNodeVal(root *TreeNode) int {
    if root.Left == nil {
        return root.Val
    }

    return getLeftestNodeVal(root.Left)
}
