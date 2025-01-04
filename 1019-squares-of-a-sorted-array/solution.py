class Solution:
    def sortedSquares(self, nums: List[int]) -> List[int]:
        res = [0] * len(nums)
        l = 0
        r = len(nums) - 1
        for i in range(len(nums)):
            left = abs(nums[l])
            right = abs(nums[r])
            if left > right:
                res[i] = left ** 2
                l += 1
            else:
                res[i] = right ** 2
                r -= 1
        
        return res[::-1]

