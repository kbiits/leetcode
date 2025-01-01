class Solution:
    def numJewelsInStones(self, jewels: str, stones: str) -> int:
        jwl = Counter(jewels)
        count = 0
        for i in stones:
            if jwl[i] > 0:
                count += 1
        return count

