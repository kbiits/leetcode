import java.util.Stack;

class Solution {
    public int maxDepth(String s) {
        if (s.length() == 0) return 0;
        Stack<Character> stack = new Stack<>();
        
        int max = 0;
        int curr = 0;
        
        for (char c : s.toCharArray()) {
            if (c == '(') {
                stack.push(c);
            } else if (c == ')') {
                int currentSize = stack.size();
                if(currentSize > max) {
                    max = currentSize;
                }
                stack.pop();
            }
        }
        
        return max;
    }
}