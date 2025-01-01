class Solution:

    # old approach
    # we don't need set because the array is sorted
    # def removeDuplicates(self, nums: List[int]) -> int:
    #     theset = set()
    #     # count = 0
    #     fast_ptr = 0
    #     slow_ptr = 0
    #     while fast_ptr < len(nums):
    #         curr = nums[fast_ptr]
    #         if curr not in theset:
    #             theset.add(curr)
    #             nums[slow_ptr] = nums[fast_ptr]
    #             fast_ptr += 1
    #             slow_ptr += 1
    #         else:
    #             fast_ptr += 1

    #     return slow_ptr

    def removeDuplicates(self, nums: List[int]) -> int:
        ptr_replace = 1
        for checking_ptr in range(1, len(nums)):
            if nums[checking_ptr] != nums[checking_ptr - 1]:
                nums[ptr_replace] = nums[checking_ptr]
                ptr_replace += 1

        return ptr_replace
                
