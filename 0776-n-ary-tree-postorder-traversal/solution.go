/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func postorder(root *Node) []int {
    nums := make([]int, 0)

    var postOrder func(*Node)
    postOrder = func(root *Node) {
        if root == nil {
            return
        }

        for _, child := range root.Children {
            postOrder(child)
        }
        nums = append(nums, root.Val)
    }

    postOrder(root)

    return nums
}
