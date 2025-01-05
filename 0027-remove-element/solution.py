class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        ptr_write = 0
        for i in range(len(nums)):
            if nums[i] != val:
                nums[ptr_write] = nums[i]
                ptr_write += 1
        return ptr_write
