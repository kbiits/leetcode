class Solution {
    public int[] buildArray(int[] arr) {
        int[] result = new int[arr.length];
        for (int i = 0; i < arr.length; i++) {
            result[i] = arr[arr[i]];
        }
        return result;
    }
}