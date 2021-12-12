class Solution {
    public List<Boolean> kidsWithCandies(int[] candies, int extraCandies) {
        int max = -1;
        for (int i = 0; i < candies.length; i++) {

            if (candies[i] > max) {
                max = candies[i];
            }
            candies[i] += extraCandies;
        }
        List<Boolean> result = new ArrayList<>();
        for (int i : candies) {
            if (i >= max) {
                result.add(true);
            } else {
                result.add(false);
            }
        }
        return result;
    }
}