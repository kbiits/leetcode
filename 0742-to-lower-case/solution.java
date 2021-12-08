class Solution {
    public String toLowerCase(String s) {
        StringBuilder sb = new StringBuilder();
        int delta = 'a' - 'A';
        for (char ch : s.toCharArray()) {
            if (ch >= 'A' && ch <= 'Z') {
                sb.append((char) (ch + delta));
            } else {
                sb.append(ch);
            }
        }
        return sb.toString();
    }
}
