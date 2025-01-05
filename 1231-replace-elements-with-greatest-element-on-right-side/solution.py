class Solution:
    def replaceElements(self, arr: List[int]) -> List[int]:
        n = len(arr)

        curr_max = arr[n - 1]
        for i in range(n - 1, -1, -1):
            temp = arr[i]
            arr[i] = curr_max
            if temp > curr_max:
                curr_max = temp

        arr[n - 1] = -1
        return arr
