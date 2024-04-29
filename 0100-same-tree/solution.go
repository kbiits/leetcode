/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSameTree(p *TreeNode, q *TreeNode) bool {
    return recurse(p, q)
}

func recurse(p *TreeNode, q *TreeNode) bool {
    if p == nil && q != nil {
        return false
    } else if p != nil && q == nil {
        return false
    } else if p == nil && q == nil {
        return true
    }

    if p.Val != q.Val {
        return false
    }

    val1 := recurse(p.Left, q.Left)
    val2 := recurse(p.Right, q.Right)

    return val1 && val2
}
