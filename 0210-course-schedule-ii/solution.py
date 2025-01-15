class Solution:
    def findOrder(self, numCourses: int, prerequisites: List[List[int]]) -> List[int]:
        
        preq_map = defaultdict(list)
        for preq in prerequisites:
            preq_map[preq[0]].append(preq[1])

        res = []
        res_set = set()
        visited_set = set()
        def dfs_course_order(course):
            if course in visited_set:
                return False

            if course in res_set:
                return True

            visited_set.add(course)

            for p in preq_map[course]:
                if not dfs_course_order(p):
                    return False

            preq_map[course] = []
            visited_set.remove(course)
            res.append(course)
            res_set.add(course)

            return True

        for course in range(numCourses):
            if not dfs_course_order(course):
                return []

        return res
            # [0, 1, 2, 3]
            # 0 => []
            # 1 => []
            # 2 => []
            # 3 => [1, 2]


