class Solution:
    def numberOfSubstrings(self, s: str, k: int) -> int:
        res = 0
        cnt = Counter()
        l = 0
        r = 0
        n = len(s)
        while r < n:
            c = s[r]
            cnt[c] += 1
            while cnt[c] >= k:
                res += n - r
                cnt[s[l]] -= 1
                l += 1
            r += 1

        return res
