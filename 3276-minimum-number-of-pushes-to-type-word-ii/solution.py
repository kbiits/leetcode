class Solution:
    def minimumPushes(self, word: str) -> int:
        sorted_word = Counter(word).most_common()
        
        res = 0
        for i in range(len(sorted_word)):
            # 0st until 7th can be pressed by 1 times
            # 8th until 15th can be pressed by 2 times

            times = 1 + (i // 8)
            res += times * sorted_word[i][1]
        
        return res

