class Solution {
    public int[] createTargetArray(int[] nums, int[] index) {
        int n = nums.length;
        int[] res = new int[n];
        for (int i = 0; i < n; i++) {
            int indexNth = index[i];
            for (int j = n - 1; j > indexNth; j--) {
                res[j] = res[j - 1];
            }
            res[indexNth] = nums[i];
        }

        return res;
    }
}
