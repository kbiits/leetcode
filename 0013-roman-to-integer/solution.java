class Solution {
    public int romanToInt(String s) {
        int res = 0;
        int max = 0;
        for (int i = s.length() - 1; i >= 0; i--) {
            int value = getIntFromCharRoman(s.charAt(i));
            if (value >= max) {
                max = value;
                res += value;
            } else
                res -= value;
        }
        return res;
    }
    
    public static int getIntFromCharRoman(char c) {
        int value = 0;
        switch (c) {
            case 'I':
                value = 1;
                break;
            case 'V':
                value = 5;
                break;
            case 'X':
                value = 10;
                break;
            case 'L':
                value = 50;
                break;
            case 'C':
                value = 100;
                break;
            case 'D':
                value = 500;
                break;
            case 'M':
                value = 1000;
                break;
            default:
                break;
        }
        return value;
    }
}
