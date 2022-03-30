import java.util.Map;
import java.util.HashMap;
import java.util.Arrays;


class Solution {
    public int[] smallerNumbersThanCurrent(int[] nums) {
        int[] counts = new int[101];
        
        for(int i = 0; i < nums.length; i++) {
            counts[nums[i]]++;
        }
        
        for(int i = 1; i < 101; i++) {
            counts[i] += counts[i - 1];
        }
        
        for(int i = 0; i < nums.length; i++) {
            int num = nums[i];
            
            nums[i] = num == 0 ? 0 : counts[nums[i] - 1];
        }
        
        return nums;
    }
}
