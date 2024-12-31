from collections import deque
class Solution:
    def findWinningPlayer(self, skills: List[int], k: int) -> int:
        champSkill, champion, tally = skills.pop(0), 0, 0

        for i, skill in enumerate(skills, start = 1):
            
            if skill > champSkill: 
                champion, champSkill, tally = i, skill, 0
                
            tally+= 1

            if tally == k:
                return champion
        
        return champion
