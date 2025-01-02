class Solution:
    def orangesRotting(self, grid: List[List[int]]) -> int:
        # find rotten fruit
        m = len(grid)
        n = len(grid[0])
        q = []

        fresh_count = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 2:
                    q.append((i, j))
                elif grid[i][j] == 1:
                    fresh_count += 1

        directions = [[0, 1], [0, -1], [1, 0], [-1, 0]]
        
        count = 0
        while len(q) > 0 and fresh_count > 0:
            for _ in range(len(q)):
                x, y = q.pop(0)
                for di, dj in directions:
                    i, j = x + di, y + dj
                    if i >= 0 and i < m and j >= 0 and j < n:
                        if grid[i][j] == 1:
                            grid[i][j] = 2
                            fresh_count -= 1
                            q.append((i, j))
                        
            count += 1
        return count if fresh_count == 0 else -1
