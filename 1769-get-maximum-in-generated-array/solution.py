class Solution:
    def getMaximumGenerated(self, n: int) -> int:
        nums = [-1] * (n + 1)
        nums[0] = 0
        max_num = 0
        if n >= 1:
            nums[1] = 1
            max_num = 1

        upper_bound = n // 2
        for i in range(0, upper_bound + 1):
            max_check = [max_num]
            if nums[i] == -1:
                nums[i] = i
                max_check.append(i)

            if 2 * i >= 2:
                if 2 * i <= n:
                    nums[i * 2] = nums[i]
                    max_check.append(nums[i])

                if 2 * i + 1 <= n:
                    max_check.append(nums[i] + nums[i + 1])
                    nums[i * 2 + 1] = nums[i] + nums[i + 1]

            # print(nums, end=" ")
            # print(max_check)
            max_num = max(max_check)

        return max_num

