class Solution:
    def checkIfExist(self, arr: List[int]) -> bool:
        n = len(arr)
        memo = set()
        # 2 * arr[i] => arr[j]
        # arr[j] / 2 = arr[i]
        for i in range(n):
            if 2 * arr[i] in memo:
                return True
            if arr[i] % 2 == 0 and arr[i] / 2 in memo:
                return True
            memo.add(arr[i])
        return False
