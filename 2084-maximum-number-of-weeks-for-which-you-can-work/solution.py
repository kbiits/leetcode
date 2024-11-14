class Solution:
    def numberOfWeeks(self, milestones: List[int]) -> int:
        if len(milestones) == 0:
            return 0
        
        milestones.sort(reverse=True)
        sum_rest = sum(milestones[1:])
        if milestones[0] <= sum_rest:
            return sum_rest + milestones[0]
        
        return 1 + (sum_rest * 2)

        

