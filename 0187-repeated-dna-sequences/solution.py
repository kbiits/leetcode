class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        table = {}
        for i in range(0, len(s) - 9):
            substr = s[i:i+10]
            if substr in table:
                table[substr] += 1
            else:
                table[substr] = 1
        
        res = []
        for k, v in table.items():
            if v > 1:
                res.append(k)
        return res
