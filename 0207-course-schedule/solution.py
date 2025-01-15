class Solution:
    def canFinish(self, numCourses: int, prerequisites: List[List[int]]) -> bool:
        preq_map = defaultdict(list)
        for preq in prerequisites:
            preq_map[preq[0]].append(preq[1])

        visited_set = set()
        def dfsCanFinish(course: int) -> bool:
            if course in visited_set:
                return False

            dep = preq_map[course]
            if not dep:
                return True
            
            visited_set.add(course)
            for depCourse in dep:
                if not dfsCanFinish(depCourse):
                    return False
            visited_set.remove(course)
            preq_map[course] = []

            return True

        # how to do cycle detection ?
        # we'll use a set to track all visited course
        # and if the course already in set, we assumed there's
        # cycle in it

        for course in range(numCourses):
            if not dfsCanFinish(course):
                return False
        return True
