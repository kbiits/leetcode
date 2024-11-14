class Solution:
    def climbStairs(self, n: int) -> int:
        # bottom up
        dp = {
            -1: 0,
            0: 1
        }
        for i in range(1, n + 1):
            dp[i] = dp[i - 1] + dp[i - 2]

        return dp[n]
