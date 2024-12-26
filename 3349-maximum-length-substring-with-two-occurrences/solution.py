class Solution:
    def maximumLengthSubstring(self, s: str) -> int:
        l = 0
        r = 1
        res = 0
        dct = {}
        dct[s[l]] = 1
        while r < len(s):
            if s[r] not in dct:
                dct[s[r]] = 1
            else:
                dct[s[r]] += 1
                while dct[s[r]] > 2:
                    dct[s[l]] -= 1
                    l += 1
            
            res = max(res, r - l + 1)
            r += 1
    
        return res
                
