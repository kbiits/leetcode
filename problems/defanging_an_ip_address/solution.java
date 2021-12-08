class Solution {
    public String defangIPaddr(String ipAddr) {
        StringBuilder sb = new StringBuilder();
        for (char ch : ipAddr.toCharArray()) {
            if (ch == '.') {
                sb.append("[.]");
            } else {
                sb.append(ch);
            }
        }
        return sb.toString();
    }
}