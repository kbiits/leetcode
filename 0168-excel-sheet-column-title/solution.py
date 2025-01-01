class Solution:
    def __init__(self):
        self.chars = [chr(ord('A') + i) for i in range(26)]
    def convertToTitle(self, n: int) -> str:
        chars = self.chars
        res = ""

        while n != 0:
            n -= 1
            mod = n % 26
            res += chars[mod]
            n //= 26


        # if n > 0:
        #     res += chars[n - 1]

        res = "".join(list(res)[::-1])
        return res
        

