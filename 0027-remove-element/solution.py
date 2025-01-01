class Solution:
    def removeElement(self, nums: List[int], val: int) -> int:
        ptr_replace = 0
        for ptr_check in range(len(nums)):
            if nums[ptr_check] != val:
                nums[ptr_replace] = nums[ptr_check]
                ptr_replace += 1

        return ptr_replace
