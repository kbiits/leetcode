class Solution {
public:
    int search(vector<int>& nums, int target) {
        return searchRecurse(nums, target, 0, nums.size() - 1);
    }
    
    int searchRecurse(vector<int>& nums, int target, int left, int right) {
        if (left > right) {
            return -1;
        }
        
        int mid = left + (right - left) / 2;
        int curr = nums[mid];
        
        if (curr == target) {
            return mid;
        } else if (target < curr) {
//             search left
            return searchRecurse(nums, target, left, mid - 1);
        } else {
//             search right
            return searchRecurse(nums, target, mid + 1, right);
        }
    }
};