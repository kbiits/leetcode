class Solution {
    public boolean arrayStringsAreEqual(String[] word1, String[] word2) {
        int pointer1Inner, pointer2Inner, pointer1Outer, pointer2Outer;
        pointer1Inner = pointer1Outer = pointer2Inner = pointer2Outer = 0;

        while (pointer1Outer < word1.length && pointer2Outer < word2.length) {
            String curr1 = word1[pointer1Outer];
            String curr2 = word2[pointer2Outer];
            int n1 = curr1.length();
            int n2 = curr2.length();
            while (pointer1Inner < n1 && pointer2Inner < n2) {
                if (curr1.charAt(pointer1Inner) != curr2.charAt(pointer2Inner)) {
                    return false;
                }

                pointer1Inner++;
                pointer2Inner++;
            }

            if (pointer1Inner >= n1) {
                pointer1Outer++;
                pointer1Inner = 0;
            }
            if (pointer2Inner >= n2) {
                pointer2Outer++;
                pointer2Inner = 0;
            }
        }

        return pointer1Outer == word1.length && pointer2Outer == word2.length;
    }
}