class Solution {
    public int reverse(int num) {
        long res = 0;
        int temp = num;
        num = num < 0 ? -num : num;
        while (num > 0) {
            res = (res * 10) + (num % 10);
            // System.out.println(res); // uncomment this line if you want to see the steps
            num = Math.floorDiv(num, 10);
        }

        if(temp < 0) {
            return (res > (Math.pow(2, 31) - 1) || res < -(Math.pow(2, 31)) ? 0 : (int) -res);
        }
        return (res > (Math.pow(2, 31) - 1) || res < -(Math.pow(2, 31)) ? 0 : (int) res);
    }
}
