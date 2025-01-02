class Solution:
    def lexicalOrder(self, n: int) -> List[int]:
        res = []
        curr = 1
        for _ in range(n):
            # 13
            # 1
            res.append(curr)
            if curr * 10 <= n:
                curr = curr * 10
            else:
                while curr % 10 == 9 or curr + 1 > n:
#                     backtrack here
                    curr //= 10
                curr += 1
        return res
                

