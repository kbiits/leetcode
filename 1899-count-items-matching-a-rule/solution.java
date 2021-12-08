class Solution {
    public int countMatches(List<List<String>> items, String ruleKey, String ruleValue) {
        int colIdx;
        switch (ruleKey) {
            case "type":
                colIdx = 0;
                break;
            case "color":
                colIdx = 1;
                break;
            default:
                colIdx = 2;
                break;
        }

        int count = 0;
        for (List<String> list : items) {
            if (list.get(colIdx).equals(ruleValue)) {
                count++;
            }
        }
        return count;
    }
}
