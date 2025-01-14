class Solution:
    def findLatestTime(self, s: str) -> str:
        s_arr = list(s)
        for i in range(len(s)):
            if s_arr[i] == "?":
                if i == 4:
                    s_arr[4] = '9'
                elif i == 3:
                    s_arr[3] = '5'
                elif i == 1:
                    if s_arr[0] == '0':
                        s_arr[1] = '9'
                    elif s_arr[0] == '1':
                        s_arr[1] = '1'
                elif i == 0:
                    if s_arr[1] == '?':
                        s_arr[0] = '1'
                    elif int(s_arr[1]) > 1:
                        s_arr[0] = '0'
                    else:
                        s_arr[0] = '1'
        return "".join(s_arr)

