class Solution:
    def isPossibleToSplit(self, nums: List[int]) -> bool:
        return max(set(Counter(nums).values())) <= 2

