class Solution:
    def arrayPairSum(self, nums: List[int]) -> int:
        nums.sort()
        left = len(nums) - 2
        right = len(nums) - 1
        
        res = 0
        while left >= 0:
            res += min(nums[left], nums[right])
            left -= 2
            right -= 2
        
        # 1, 2, 3, 4

        return res
        
