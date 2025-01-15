class Solution:
    def nthUglyNumber(self, n: int) -> int:
        primes = [2, 3, 5]
        heap = [1]
        visited = set()
        visited.add(1)

        for _ in range(n):
            prev_ugly = heappop(heap)
            for prime in primes:
                new_ugly = prev_ugly * prime
                if new_ugly not in visited:
                    heappush(heap, new_ugly)
                    visited.add(new_ugly)

        return prev_ugly
