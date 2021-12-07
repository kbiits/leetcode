class Solution {
    public int search(int[] nums, int target) {
        return binarySearchRecursion(nums, target, 0, nums.length - 1);
    }
    
    public static int binarySearchRecursion(int[] arr, int needle, int left, int right) {
        if (left > right) {
            return -1;
        }
        int mid = left + (right - left) / 2;
        if (arr[mid] == needle) {
            return mid;
        } else if (needle < arr[mid]) {
            return binarySearchRecursion(arr, needle, left, mid - 1);
        } else {
            return binarySearchRecursion(arr, needle, mid + 1, right);
        }
    }
}
