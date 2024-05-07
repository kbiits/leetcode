/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isSymmetric(root *TreeNode) bool {
    if root == nil{
        return true
    }
    q1 := []*TreeNode{root.Left}
    q2 := []*TreeNode{root.Right}

    for len(q1) > 0 && len(q2) > 0{
        currentNode1 := q1[0]
        currentNode2 := q2[0]
        q1 = q1[1:]
        q2 = q2[1:]

        if currentNode1 == nil && currentNode2 == nil {
            continue
        }

        if currentNode1 == nil || currentNode2 == nil || currentNode2.Val != currentNode1.Val {
            return false
        }

        q1 = append(q1, currentNode1.Left, currentNode1.Right)
        q2 = append(q2, currentNode2.Right, currentNode2.Left)

    }

    return true
}

