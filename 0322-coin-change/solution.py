
class Solution:
    def coinChange(self, coins: List[int], amount: int) -> int:
        dp = {
            0: 0
        }

        for sub_amount in range(1, amount + 1):
            best_min = float('inf')
            for coin in coins:
                if coin <= sub_amount:
                    best_min = min(best_min, dp[sub_amount - coin] + 1)

            dp[sub_amount] = best_min

        return dp[amount] if dp[amount] != float('inf') else -1

