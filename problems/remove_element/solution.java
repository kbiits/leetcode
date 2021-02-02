class Solution {
    public int removeElement(int[] arr, int val) {
        int n = arr.length;
        int i = 0;
        while(i < n) {
            if(arr[i] == val) {
                arr[i] = arr[n-1];
                n--;
            } else {
                i++;
            }
        }
        return n;
    }
}