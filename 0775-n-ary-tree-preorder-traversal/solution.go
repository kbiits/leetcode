/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
    nums := make([]int, 0, 50)
    var preOrder func(*Node)
    preOrder = func(root *Node) {
        if root == nil {
            return
        }

        nums = append(nums, root.Val)
        for _, child := range root.Children {
            preOrder(child)
        }
    }

    preOrder(root)

    return nums
}
