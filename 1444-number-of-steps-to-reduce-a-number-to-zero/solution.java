class Solution {
    public int numberOfSteps(int num) {
        if (num == 0)
            return 0;
        return 1 + numberOfSteps(num % 2 == 0 ? num / 2 : num - 1);
    }
}
