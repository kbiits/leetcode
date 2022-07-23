class Solution {
public:
    int searchInsert(vector<int>& nums, int target) {
        return binarySearch(nums, target, 0, nums.size() - 1);
    }
    
    int binarySearch(vector<int>& nums, int target, int left, int right) {
        
        if (left > right) {
            return left;
        }
        
        
        int mid = left + (right - left) / 2;
        int curr = nums[mid];
        
        if (curr == target) {
            return mid;
        } else if (target > curr) {
//             go to right
            return binarySearch(nums, target, mid + 1, right);
        } else {
//             go to left
            return binarySearch(nums, target, left, mid - 1);
        }
        
    }
};