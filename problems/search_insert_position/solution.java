class Solution {
    public int searchInsert(int[] arr, int target) {
        return search(arr, 0, arr.length-1, target);
    }
    
    public int search(int[] arr, int start, int end, int target) {
        if(end >= start) {
            int mid = ((end-start)/ 2) + start;
            if(arr[mid] == target) {
                return mid;
            } else if (target < arr[mid]){
                return search(arr, start, mid-1, target);
            } else {
                return search(arr, mid+1, end, target);
            }
        } else {
            return start;
        }
    }
}