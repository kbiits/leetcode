/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func constructMaximumBinaryTree(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    maxIdx := 0
    maxEl := -1
    for i, v := range nums {
        if v > maxEl {
            maxEl = v
            maxIdx = i
        }
    }

    root := &TreeNode{Val: maxEl}
    root.Left = constructMaximumBinaryTree(nums[:maxIdx])
    if maxIdx + 1 < len(nums) {
        root.Right = constructMaximumBinaryTree(nums[maxIdx+1:])
    }

    return root
}
