class Solution {
    public int countKDifference(int[] nums, int k) {
        Map<Integer, Integer> map = new HashMap<>();

        int sum = 0;
        int n = nums.length;
        for (int i = 0; i < n; i++) {
            int curr = nums[i];
            if (map.containsKey(curr + k)) {
                sum += map.get(curr + k);
            }
            if (map.containsKey(curr - k)) {
                sum += map.get(curr - k);
            }

            map.put(curr, map.getOrDefault(curr, 0) + 1);
        }

        return sum;
    }
}