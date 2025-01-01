class Solution:
    def romanToInt(self, s: str) -> int:
        table = {
            'I': 1,
            'V': 5,
            'X': 10,
            'L': 50,
            'C': 100,
            'D': 500,
            'M': 1000,
        }

        i = 0
        n = len(s)
        res = 0

        # IX => 9
        # VII => 8
        # IV => 4
        while i < n:
            c = s[i]
            if i < n - 1:
                if table[c] < table[s[i + 1]]:
                    res += table[s[i + 1]] - table[c]
                    i += 2
                    continue

            res += table[c]
            i += 1
        return res

