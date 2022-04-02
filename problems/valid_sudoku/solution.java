import java.util.HashSet;
import java.util.Set;

class Solution {
    public boolean isValidSudoku(char[][] board) {
        for (int i = 0; i < 9; i++) {
            for (int j = 0; j < 9; j++) {
                if (board[i][j] == '.') {
                    continue;
                }

                if (!isSafe(board, i, j)) {
                    return false;
                }
            }
        }

        return true;
    }
    
    private static boolean isSafe(char[][] board, int row, int col) {
        char val = board[row][col];

        // check horizontal
        for (int currCol = 0; currCol < 9; currCol++) {
            if (currCol != col && board[row][currCol] == val) {
                return false;
            }
        }

        // check vertical
        for (int currRow = 0; currRow < 9; currRow++) {
            if (currRow != row && board[currRow][col] == val) {
                return false;
            }
        }

        // check in same box
        int boxRowStartIdx = (row / 3) * 3;
        int boxColStartIdx = (col / 3) * 3;
        for (int i = 0; i < 3; i++) {
            for (int j = 0; j < 3; j++) {
                int tempRow = boxRowStartIdx + i;
                int tempCol = boxColStartIdx + j;
                if (row != tempRow && col != tempCol && board[tempRow][tempCol] == val) {
                    return false;
                }
            }
        }

        return true;
    }
}