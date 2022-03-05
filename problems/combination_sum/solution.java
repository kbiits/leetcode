class Solution {
    public List<List<Integer>> combinationSum(int[] candidates, int target) {
        List<List<Integer>> result = new ArrayList<>();
        List<Integer> currentCombination = new ArrayList<>();
        getCombinations(candidates, target, currentCombination, 0, 0, result);

        return result;
    }
    
    private static void getCombinations(int[] candidates, int target, List<Integer> currentCombination, int sum,
            int idx,
            List<List<Integer>> result) {
        if (sum > target)
            return; // backtrack
        if (sum == target) {
            List<Integer> copyList = List.copyOf(currentCombination);
            result.add(copyList);
            return;
        }

        for (int i = idx; i < candidates.length; i++) {
            currentCombination.add(candidates[i]);
            sum += candidates[i];
            getCombinations(candidates, target, currentCombination, sum, i, result);
            sum -= candidates[i];
            currentCombination.remove(currentCombination.size() - 1);
        }
    }
}