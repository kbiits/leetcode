import java.util.*;

class Solution {
    public int[] twoSum(int[] arr, int target) {
        Map<Integer, Integer> map = new HashMap<>();

        for (int i = 0; i < arr.length; i++) {
            if (map.containsKey(arr[i])) {
                return new int[] { map.get(arr[i]), i };
            }

            map.put(target - arr[i], i);
        }

        return new int[] {};
        
    }
}
