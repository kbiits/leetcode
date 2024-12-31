class Solution:
    def minElement(self, nums: List[int]) -> int:
        def sumDigit(x):
            res = 0
            while x != 0:
                last = x % 10
                res += last
                x = x // 10
            return res
        
        return min(map(sumDigit, nums))
