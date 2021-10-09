class Solution {
    public int numIdenticalPairs(int[] arr) {
     int n = 0;
        Map<Integer, Integer> map = new HashMap<>();
        for (int i : arr) {
            if (map.containsKey(i)) {
                map.put(i, map.get(i) + 1);
                n += map.get(i);
            } else {
                map.put(i, 0);
            }
        }
        return n;
    }
}
