class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        count = 1
        replace_ptr = 1
        for scan_ptr in range(1, len(nums)):
            if nums[scan_ptr] == nums[scan_ptr - 1]:
                count += 1
            else:
                count = 1

            if count <= 2:
                nums[replace_ptr] = nums[scan_ptr]
                replace_ptr += 1
        return replace_ptr

