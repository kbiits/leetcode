class Solution {
    public int[] decompressRLElist(int[] nums) {
        int size = 0;
        for (int i = 0; i < nums.length; i += 2) {
            size += nums[i];
        }
        int[] res = new int[size];
        int pointer = 0;
        for (int i = 0; i < nums.length; i += 2) {
            int freq = nums[i];
            Arrays.fill(res, pointer, pointer + freq, nums[i + 1]);
            pointer += freq;
        }

        return res;
    }
}