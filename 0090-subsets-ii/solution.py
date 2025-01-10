class Solution:
    def subsetsWithDup(self, nums: List[int]) -> List[List[int]]:
        result = []
        def solve(idx, res):
            result.append(res.copy())

            for i in range(idx, len(nums)):
                if i != idx and nums[i] == nums[i - 1]:
                    continue
                res.append(nums[i])
                solve(i + 1, res)
                res.pop()

        nums.sort()
        solve(0, [])
        return result
