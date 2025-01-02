class Solution:
    def isValidSudoku(self, board: List[List[str]]) -> bool:
        board_table = {i: set() for i in range(9)}
        col_table = {i: set() for i in range(9)}
        for i in range(9):
            row = set()
            for j in range(9):
                cur = board[i][j]
                if cur == ".":
                    continue

                if cur in row:
                    return False
                else:
                    row.add(cur)

                col = j
                if cur in col_table[col]:
                    return False
                else:
                    col_table[col].add(cur)
                
                board_num = ((i // 3)*3) + (j // 3)
                if cur in board_table[board_num]:
                    return False
                else:
                    board_table[board_num].add(cur)

        return True

