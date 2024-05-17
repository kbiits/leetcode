/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }

    maxValDepth := 0
    for _, child := range root.Children {
        childDepth := maxDepth(child)
        if childDepth > maxValDepth {
            maxValDepth = childDepth
        }
    }

    return 1 + maxValDepth
}
