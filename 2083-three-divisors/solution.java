class Solution {
    public boolean isThree(int n) {
        int divisor = 2;
        if(n < 2) {
            return false;
        }
        for(int i = 2; i < n; i++) {
            if(n % i == 0) divisor++;
        }
        
        return divisor == 3;
    }
}
