class Solution {
    public int removeDuplicates(int[] nums) {
        int numOfDuplicates = 1;
        for(int i = 1; i < nums.length; i++) { // O(n) solution 
            if(nums[i] != nums[i-1])
                nums[numOfDuplicates++] = nums[i];
        }
        return numOfDuplicates;
    }
}
