class Solution {
    public void solveSudoku(char[][] board) {
        char[][] problems = new char[9][9];
        for (int i = 0; i < 9; i++) {
            problems[i] = Arrays.copyOf(board[i], 9);
        }
        solve(problems, board, 0, 0);
    }
    
    static void solve(char[][] board, char[][] result, int row, int col) {
        if (row >= 9) {
            for (int i = 0; i < 9; i++) {
                for (int j = 0; j < 9; j++) {
                    result[i][j] = board[i][j];
                }
            }
            return;
        }

        if (col >= 9) {
            solve(board, result, row + 1, 0);
            return;
        }

        // skip if there's a number in the cell
        if (board[row][col] != '.') {
            solve(board, result, row, col + 1);
            return;
        }

        for (int i = 1; i <= 9; i++) {
            char currChar = (char) (i + '0');
            if (isSafe(currChar, row, col, board)) {
                // mark visited
                board[row][col] = currChar;
                solve(board, result, row, col + 1);
                // unmark visited
                board[row][col] = '.';
            }
        }
    }

    static boolean isSafe(char val, int row, int col, char[][] board) {
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
                if (board[boxRowStartIdx + i][boxColStartIdx + j] == val) {
                    return false;
                }
            }
        }

        return true;
    }
}