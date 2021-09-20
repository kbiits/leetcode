class Solution {
    public int singleNumber(int[] nums) {
        int xorVal = 0;
        for(int i : nums) {
            xorVal ^= i;
        }
        
        return xorVal;
    }
}
