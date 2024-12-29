class Solution:
    def maxScore(self, n: int, k: int, stayScore: List[List[int]], travelScore: List[List[int]]) -> int:
        dp = [[0] * n for i in range(k + 1)]
        # dp[day][city]
        # stayScore[day][city]
        # travelScore[curr_city][dest_city]

        for city in range(n):
            dp[1][city] = stayScore[0][city]
            for city2 in range(n): # city2 => from
                dp[1][city] = max(dp[1][city], travelScore[city2][city])

        for i in range(2, k+1):
            for j in range(n):
                for k in range(n):
                    if k == j:
                        dp[i][j] = max(dp[i][j], dp[i-1][j] + stayScore[i - 1][j])
                    else:
                        dp[i][j] = max(dp[i][j], dp[i-1][k] + travelScore[k][j])

        return max(dp[-1])
        

