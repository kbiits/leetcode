class Solution {
    public int subtractProductAndSum(int n) {
        int products = 1;
        int sum = 0;

        while (n > 0) {
            int temp = n % 10;
            products *= temp;
            sum += temp;
            n = n / 10;
        }
        
        return products - sum;
    }
}