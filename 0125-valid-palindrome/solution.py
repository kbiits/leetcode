class Solution:
    def isPalindrome(self, s: str) -> bool:
        l = 0
        r = len(s) - 1
        s = s.lower()
        def is_alphanum(c):
            return (ord('a') <= ord(c) <= ord('z')) or (ord('0') <= ord(c) <= ord('9'))

        while l <= r:
            if not is_alphanum(s[l]):
                l += 1
                continue
            if not is_alphanum(s[r]):
                r -= 1
                continue
            
            if s[l] != s[r]:
                return False
            
            l += 1
            r -= 1

        return True
