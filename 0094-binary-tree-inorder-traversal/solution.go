/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderTraversal(root *TreeNode) []int {
    res := []int{}
    recurse(&res, root)

    return res
}

func recurse(res *[]int, root *TreeNode) {
    if root == nil {
        return
    }

    recurse(res, root.Left)
    *res = append(*res, root.Val)
    recurse(res, root.Right)

    return
}
