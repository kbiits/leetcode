class Solution {
    public int sumOddLengthSubarrays(int[] arr) {
        int sum = 0;
        int n = arr.length;
        for (int i = 0; i < n; i++) {
            // (x + 1) / 2, because we just want to add odd length subarray
            sum += ((i + 1) * (n - i) + 1) / 2 * arr[i];
        }

        return sum;
    }
}
