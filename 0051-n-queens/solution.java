import java.util.List;
import java.util.ArrayList;


class Solution {
    public List<List<String>> solveNQueens(int n) {
        List<List<String>> result = new ArrayList<>();
        nQueens(result, new boolean[n][n], 0);
        return result;
    }
    
    private static void nQueens(List<List<String>> result, boolean[][] board, int row) {
        int n = board.length;
        if (row == n) {
            List<String> res = convertBoolArrToListArr(board);
            result.add(res);
            return;
        }

        for (int i = 0; i < n; i++) {
            if (isSafe(row, i, board)) {
                board[row][i] = true;
                nQueens(result, board, row + 1);
                board[row][i] = false;
            }
        }
    }

    private static List<String> convertBoolArrToListArr(boolean[][] board) {
        List<String> res = new ArrayList<>();

        for (boolean[] row : board) {
            String temp = new String("");
            for (boolean col : row) {
                if (col) {
                    temp += "Q";
                } else {
                    temp += ".";
                }
            }
            res.add(temp);
        }

        return res;
    }

    private static boolean isSafe(int row, int col, boolean[][] board) {
        int nIdx = board.length - 1;

        // check top
        for (int i = 0; i < row; i++) {
            if (board[i][col]) {
                return false;
            }
        }

        // check top left
        int minLeft = Math.min(row, col);
        for (int i = minLeft; i >= 0; i--) {
            if (board[row - i][col - i]) {
                return false;
            }
        }

        // check top right
        int minRight = Math.min(row, nIdx - col);
        for (int i = minRight; i > 0; i--) {
            if (board[row - i][col + i]) {
                return false;
            }
        }

        return true;
    }
}
