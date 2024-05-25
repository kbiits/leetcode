/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func findTarget(root *TreeNode, k int) bool {
    var mapComplement = map[int]bool{}
    var search func(root *TreeNode) bool
	search = func(root *TreeNode) bool {
        if root == nil {
            return false
        }

        if mapComplement[k - root.Val] {
            return true
        }

        mapComplement[root.Val] = true

        left := search(root.Left)
        right := search(root.Right)
        return left || right
    }
    return search(root)
}

