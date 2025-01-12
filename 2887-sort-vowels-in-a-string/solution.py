class Solution:
    def sortVowels(self, s: str) -> str:
        vowels = set(['A', 'I', 'U', 'E', 'O', 'a', 'i', 'u', 'e', 'o'])
        temp = s
        res = []
        vowels_arr = []
        for i in range(len(s)):
            c = s[i]
            if c in vowels:
                vowels_arr.append(c)
                res.append("*")
            else:
                res.append(c)

        vowels_arr.sort()
        ptr = 0
        for i in range(len(res)):
            c = res[i]
            if c == "*":
                res[i] = vowels_arr[ptr]
            else:
                continue
            ptr += 1

        return "".join(res)


