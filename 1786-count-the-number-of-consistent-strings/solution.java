class Solution {
    public int countConsistentStrings(String allowed, String[] words) {
        Set<Character> map = new HashSet<>();
        for (char i : allowed.toCharArray()) {
            map.add(i);
        }

        int result = words.length;
        for (String word : words) {
            for (char a : word.toCharArray()) {
                if (!map.contains(a)) {
                    result--;
                    break;
                }
            }
        }

        return result;
    }
}
