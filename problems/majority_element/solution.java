class Solution {
    public int majorityElement(int[] nums) {
        
        int major = nums[0]; // assume major is nums[0]
        int counter = 1; // bcs we assume the major is nums[0] so counter should be 1
        for(int i = 1; i<nums.length; i++) {
            if(counter == 0) {
//                 change the major element
                major = nums[i];
                counter++;
            } else if(nums[i] == major) {
                counter++;
            } else { // nums[i] not same with major
                counter--;
            }
        }
        return major;
    }
}