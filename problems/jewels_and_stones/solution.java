class Solution {
    public int numJewelsInStones(String jewels, String stones) {
        // counting sort like implementation
        int[] cnt = new int[128];
        for (char c : stones.toCharArray())
            cnt[c]++;
        int ans = 0;
        for (char c : jewels.toCharArray())
            ans += cnt[c];
        return ans;
    }
}