class Solution {
    public String removeOuterParentheses(String s) {
        int counter = 0;
        String res = "";
        for (char c : s.toCharArray()) {
            if(c == '(') {
                if(counter++ != 0) {
                    res += "(";
                }
            } else {
                if(--counter != 0) {
                    res += ")";
                }
            }
        }
        
        return res;
        
    }
}