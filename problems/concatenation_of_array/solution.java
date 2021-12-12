class Solution {
    public int[] getConcatenation(int[] arr) {
        int[] res = new int[arr.length * 2];
        for (int i = 0; i < arr.length; i++) {
            res[i] = arr[i];
            res[i + arr.length] = arr[i];
        }
        return res;
    }
}