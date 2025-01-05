import heapq
class Solution:
    def kWeakestRows(self, mat: List[List[int]], k: int) -> List[int]:
        q = []
        m = len(mat)
        n = len(mat[0])
        for i in range(m):
            score = mat[i].count(1)
            heapq.heappush(q, (-score, -i))
            if len(q) > k:
                heapq.heappop(q)
        
        res = [-heapq.heappop(q)[1] for i in range(k)]
        return res[::-1]


        

