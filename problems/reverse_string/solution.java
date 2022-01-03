class Solution {
    public void reverseString(char[] s) {
        reverse(s, 0, s.length - 1);
    }
    
    public char[] reverse(char[] s, int idxLeft, int idxRight) {
        if(idxLeft >= idxRight) {
            return s;
        }
        char temp = s[idxLeft];
        s[idxLeft] = s[idxRight];
        s[idxRight] = temp;
        
        return reverse(s, idxLeft + 1, idxRight - 1);
    }
}