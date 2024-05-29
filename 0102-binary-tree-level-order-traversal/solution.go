import "container/list"
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func levelOrder(root *TreeNode) [][]int {
    nodes := [][]int{}
    if root == nil {
        return nodes
    }
    
    queue := list.New()
    queue.PushBack(root)

    for queue.Len() > 0 {
        n := queue.Len()
        levelArr := make([]int, n)
        for i := 0; i < n; i++ {
            // pop
            el := queue.Front()
            queue.Remove(el)
            node := el.Value.(*TreeNode)
            levelArr[i] = node.Val
            if node.Left != nil {
                queue.PushBack(node.Left)
            }
            if node.Right != nil {
                queue.PushBack(node.Right)
            }
        }
        nodes = append(nodes, levelArr)
    }

    return nodes
}
