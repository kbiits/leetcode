class Solution {
    public String freqAlphabets(String s) {
        StringBuilder sb = new StringBuilder();
        int delta1 = 'a' - '1';
        for (int i = 0; i < s.length(); i++) {
            if ((i + 2) >= s.length() || s.charAt(i + 2) != '#') {
                sb.append((char) (s.charAt(i) + delta1));
            } else {
                int intValue = Integer.parseInt(s.substring(i, i + 2));
                sb.append((char) ('a' - 1 + intValue));
                i += 2;
            }
        }
        return sb.toString();
    }
}