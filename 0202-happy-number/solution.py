class Solution:
    def isHappy(self, n: int) -> bool:
        def calc(n):
            res = 0
            while n != 0:
                last = n % 10
                res += last**2
                n = n // 10
            return res

        check = set()
        while n != 1:
            temp = calc(n)
            if temp in check:
                return False

            check.add(temp)
            n = temp

        return True


            
