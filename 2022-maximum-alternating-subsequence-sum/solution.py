class Solution:
    nums: List[int] = []
    def maxAlternatingSum(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0

        dp = {}
        def helper(i, isEven) -> int:
            if i == len(nums):
                return 0

            if (i, isEven) in dp:
                return dp[(i, isEven)]

            cur = -nums[i] if isEven else nums[i]

            # set the dp
            dp[(i, isEven)] = max(
                cur + helper(i + 1, not isEven), # not skip
                helper(i + 1, isEven) # skip, keep the flag because it's reindexed
            )
            return dp[(i, isEven)]

        val = helper(0, False)
        return val

    

