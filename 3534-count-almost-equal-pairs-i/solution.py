class Solution:
    def countPairs(self, nums: List[int]) -> int:

        @cache
        def is_almost(x, y):
            x = str(x)
            y = str(y)
            max_len = max(len(x), len(y))
            while len(x) < max_len:
                x = "0" + x
            while len(y) < max_len:
                y = "0" + y

            diff = 0
            x_map = Counter()
            y_map = Counter()
            for i in range(max_len):
                if x[i] != y[i]: # order matters
                    diff += 1
                x_map[x[i]] += 1
                y_map[y[i]] += 1

            return diff <= 2 and x_map == y_map # comparing counter because ordered not matters
        
        count = 0
        for i in range(len(nums)):
            for j in range(i + 1, len(nums)):
                x = nums[i]
                y = nums[j]
                if x == y:
                    count += 1
                    continue

                if is_almost(x, y):
                    count += 1
        return count

