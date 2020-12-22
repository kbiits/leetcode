class Solution {
    public boolean isPalindrome(int x) {
        if(x >= 0) {
            int y = 0;
            int temp = x;
            while(x > 0) {
                y = (y*10) + (x%10);
                x /= 10;
            }
            return temp == y;
        }
        return false;
    }
}