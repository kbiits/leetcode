class Solution:
    def sortArrayByParity(self, nums: List[int]) -> List[int]:
        n = len(nums)
        l = 0 # l used to replace/write
        r = n - 1 # r used to check/read
        while l < r:
            if nums[r] % 2 == 0:
                nums[l], nums[r] = nums[r], nums[l]
                l += 1
            else:
                r -= 1

        return nums
