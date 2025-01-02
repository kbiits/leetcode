# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isBalanced(self, root: Optional[TreeNode]) -> bool:
        def helper(root):
            if root is None:
                return (True, 0)
            
            leftDepth = helper(root.left)
            if leftDepth[0] is False:
                return (False, leftDepth[1])

            rightDepth = helper(root.right)
            if rightDepth[0] is False:
                return (False, rightDepth[1])

            if abs(leftDepth[1] - rightDepth[1]) > 1:
                return (False, 0)

            return (True, 1 + max(leftDepth[1], rightDepth[1]))

        return helper(root)[0]
        
