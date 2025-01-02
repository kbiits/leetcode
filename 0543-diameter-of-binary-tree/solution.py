# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def diameterOfBinaryTree(self, root: Optional[TreeNode]) -> int:
        max_value = -1

        def dfs(root):
            nonlocal max_value
            if root is None:
                return 0

            left = dfs(root.left)
            right = dfs(root.right)
        
            if left + right > max_value:
                max_value = left + right

            return 1 + max(left, right)

        dfs(root)
        return max_value
