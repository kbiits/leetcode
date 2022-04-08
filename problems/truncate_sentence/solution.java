class Solution {
    public String truncateSentence(String s, int k) {
        StringBuilder res = new StringBuilder();
        for (char c : s.toCharArray()) {
            if (c == ' ') {
                k--;
            }
            if (k == 0) {
                break;
            }
            res.append(c);
        }

        return res.toString();
    }
}