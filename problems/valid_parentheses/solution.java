import java.util.Stack;

class Solution {
    public boolean isValid(String s) {
        Stack<Character> stack = new Stack<>();
        
        for (int i = 0; i < s.length(); i++) {
            Character curr = s.charAt(i);
            if(isOpenParenthese(curr)) {
                stack.push(curr);
            } else {
                
                if(stack.isEmpty()) {
                    return false;
                } else if (!checkValidParenthese(stack.peek(), curr)) {
                    return false;
                } else {
                    stack.pop();                    
                }
                
            }
        }

        return stack.isEmpty();
    }
               
    public boolean checkValidParenthese(char open, char close) {
        if(open == '(' && close == ')') return true;
        if(open == '[' && close == ']') return true;
        if(open == '{' && close == '}') return true;
        
        return false;
    }
    
    public boolean isOpenParenthese(char ch) {
        return ch == '(' || ch == '{' || ch == '[';
    }
}