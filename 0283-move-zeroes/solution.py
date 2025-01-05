class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        ptr_write = 0
        n = len(nums)
        count_zero = 0
        for i in range(n):
            if nums[i] == 0:
                count_zero += 1
            else:
                nums[ptr_write] = nums[i]
                ptr_write += 1
        
        for i in range(n - 1, n - 1 - count_zero, -1):
            nums[i] = 0
        
