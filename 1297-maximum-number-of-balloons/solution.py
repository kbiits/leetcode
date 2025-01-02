class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        cnt = Counter(text)
        min_val = float('inf')
        for i, need in [('b', 1), ('a', 1), ('l', 2), ('o', 2), ('n', 1)]:
            min_val = min(min_val, cnt[i] // need)

        return min_val
