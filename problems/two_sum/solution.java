class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> maps = new HashMap();
        for(int i = 0; i < nums.length; maps.put(nums[i], i++)) {
            if(maps.containsKey(target-nums[i]))
                return new int[] {maps.get(target-nums[i]), i};
        }
        return new int[2];
    }
}