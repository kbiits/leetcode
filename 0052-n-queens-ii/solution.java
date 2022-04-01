class Solution {
    public int totalNQueens(int n) {
        return nQueens(new boolean[n][n], 0);
    }
    
    private static int nQueens(boolean[][] board, int row) {
        int n = board.length;
        if (row == n) {
            return 1;
        }

        int count = 0;
        for (int i = 0; i < n; i++) {
            if (isSafe(board, row, i)) {
                board[row][i] = true;
                count += nQueens(board, row + 1);
                board[row][i] = false;
            }
        }

        return count;
    }

    private static boolean isSafe(boolean[][] board, int row, int col) {
        int nIdx = board.length - 1;

        // check top
        for (int i = 0; i < row; i++) {
            if (board[i][col]) {
                return false;
            }
        }

        // check top left
        int minLeft = Math.min(row, col);
        for (int i = minLeft; i > 0; i--) {
            if (board[row - i][col - i]) {
                return false;
            }
        }

        int minRight = Math.min(nIdx - col, row);
        for (int i = minRight; i > 0; i--) {
            if (board[row - i][col + i]) {
                return false;
            }
        }

        return true;
    }
}
