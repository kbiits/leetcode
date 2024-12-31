import heapq

class Solution:
    def resultsArray(self, queries: List[List[int]], k: int) -> List[int]:
        pq = []
        res = []
        for q in queries:
            val = abs(q[0]) + abs(q[1])
            # -val means invert the value
            # because python heapq has no max-heap implementation, and we want to use max heap
            # we only store k elements in max-heap
            # so, the 0 index in max-heap should be the kth nearest to origin

            heapq.heappush(pq, -val)
            if len(pq) > k:
                heapq.heappop(pq)
            if len(pq) == k:
                res.append(-pq[0])
            else:
                res.append(-1)
        return res
