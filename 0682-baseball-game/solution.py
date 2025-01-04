class Solution:
    def calPoints(self, operations: List[str]) -> int:
        # just keep 2 last element
        res = []

        res_int = 0
        for op in operations:
            if op == "C":
                s = res.pop()
                res_int -= s
            elif op == "D":
                res.append(res[-1] * 2)
                res_int += res[-1]
            elif op == "+":
                res.append(res[-1] + res[-2])
                res_int += res[-1]
            else:
                res.append(int(op))
                res_int += res[-1]

        return res_int

