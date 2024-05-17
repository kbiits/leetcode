/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func averageOfLevels(root *TreeNode) []float64 {
    result := []float64{}
    queue := []*TreeNode{root}

    for len(queue) > 0 {
        currLevelCount := len(queue)
        currLevelSum := 0

        // processing per level (root => 0, root childs => 2, root child childs => 4)
        for i := 0; i < currLevelCount; i++ {
            currNode := queue[0]
            // pop
            queue = queue[1:]
            if currNode.Left != nil {
                queue = append(queue, currNode.Left)
            }
            if currNode.Right != nil {
                queue = append(queue, currNode.Right)
            }

            currLevelSum += currNode.Val
        }

        result = append(result, float64(currLevelSum) / float64(currLevelCount))
    }

    return result
}
