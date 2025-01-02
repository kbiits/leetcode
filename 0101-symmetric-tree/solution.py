# Definition for a binary tree node.
# class TreeNode:
#     def __init__(self, val=0, left=None, right=None):
#         self.val = val
#         self.left = left
#         self.right = right
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        def helper(p, q):
            if p is None:
                if q is None:
                    return True
                if q is not None:
                    return False
            if q is None:
                if p is None:
                    return True
                if p is not None:
                    return False

            if p.val != q.val:
                return False

            a = helper(p.left, q.right)
            b = helper(p.right, q.left)
            return a and b

        if root.left is None:
            return root.right is None
        if root.right is None:
            return root.left is None

        return helper(root.left, root.right)
        
