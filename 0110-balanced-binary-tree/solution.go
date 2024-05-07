/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    return getHeight(root) >= 0
}

func getHeight(root *TreeNode) int {
    if root == nil {
        return 0
    }

    left := getHeight(root.Left)
    right := getHeight(root.Right)
    if left == -1 || right == -1 {
        return -1
    }

    diff := absDiff(left, right)
    if diff > 1 {
        return -1
    }

    max := left
    if right > max {
        max = right
    }


    return max + 1
}

func absDiff(a, b int) int {
    if a < b {
        return b - a
    }

    return a - b
}

//            1
//         2     2
//       3         3
//     4             4
