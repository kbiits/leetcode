class Solution:
    def restoreMatrix(self, rowSum: List[int], colSum: List[int]) -> List[List[int]]:
        # 3
        # 8

        # 4 7

        # m = 2
        # n = 2
        m = len(rowSum)
        n = len(colSum)
        matrix = []
        for i in range(m):
            matrix.append([])
            for j in range(n):
                el = min(rowSum[i], colSum[j])
                rowSum[i] -= el
                colSum[j] -= el

                matrix[i].append(
                    el
                )
        
        return matrix
        
