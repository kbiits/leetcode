class Solution:
    def findClosestNumber(self, nums: List[int]) -> int:
        closest = nums[0]
        abs_val = closest - 0
        if abs_val < 0:
            abs_val = -1 * abs_val

        for i in range(1, len(nums)):
            n = nums[i]
            diff = n - 0
            if diff < 0:
                diff = -1 * diff

            if diff < abs_val or (diff == abs_val and n > closest):
                closest = n
                abs_val = diff
            
        return closest
