class Solution:
    def validMountainArray(self, arr: List[int]) -> bool:
        if len(arr) < 3:
            return False

        highest_point = 0
        for i in range(1, len(arr)):
            if arr[i] > arr[highest_point]:
                highest_point = i

        increasing_exists = False
        for i in range(highest_point):
            if arr[i] < arr[i + 1]:
                increasing_exists = True
            else:
                return False

        decreasing_exists = False
        for i in range(highest_point, len(arr) - 1):
            if arr[i] > arr[i + 1]:
                decreasing_exists = True
            else:
                return False
        return increasing_exists and decreasing_exists
