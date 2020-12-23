class Solution {
    public boolean isValid(String str) {
       Stack<Character> charStack = new Stack<>();
        for (int i = 0; i < str.length(); i++) {
            char c = str.charAt(i);
            if (c == '(' || c == '{' || c == '[') { 
                charStack.push(c);
            } else {
                if (!charStack.isEmpty() && (charStack.peek() == c - 1 || charStack.peek() == c - 2)) {
                    charStack.pop();
                } else
                    return false;
            }
        }
        if (charStack.isEmpty())
            return true;
        else
            return false;
    }
}