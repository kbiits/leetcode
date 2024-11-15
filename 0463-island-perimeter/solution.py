class Solution:
    def islandPerimeter(self, grid: List[List[int]]) -> int:
        def calc_perimeter(i, j):
            res = 0

            # check up border
            if i == 0 or grid[i - 1][j] == 0:
                res += 1
            # check bottom border
            if i == len(grid) - 1 or grid[i + 1][j] == 0:
                res += 1

            # check left border
            if j == 0 or grid[i][j - 1] == 0:
                res += 1
            # check right border
            if j == len(grid[0]) - 1 or grid[i][j + 1] == 0:
                res += 1

            return res

        res = 0
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                if grid[i][j] == 1:
                    res += calc_perimeter(i, j)

        return res
            


