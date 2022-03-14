class Solution {
    public int maximumWealth(int[][] accounts) {
        int max = 0;
        
        for (int[] i : accounts) {
            int temp = 0;
            for (int j : i) {
                temp += j;
            }
            if(temp > max) max = temp;
        }
        return max;
    }
}