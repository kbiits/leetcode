class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        n = len(heights)
        counting_sort = [0] * (max(heights) + 1)
        for height in heights:
            counting_sort[height] += 1
        
        correct_data = []
        for i, count in enumerate(counting_sort):
            if count > 0:
                for j in range(count):
                    correct_data.append(i)

        count = 0
        for i in range(n):
            if correct_data[i] != heights[i]:
                count += 1
        return count
        
