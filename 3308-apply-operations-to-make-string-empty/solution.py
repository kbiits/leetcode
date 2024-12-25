class Solution:
    def lastNonEmptyString(self, s: str) -> str:
        counter = {}
        for i in range(len(s)):
            c = s[i]
            if not c in counter:
                counter[c] = {"idx": i, "val": 1}
            else:
                prev = counter[c]
                prev["idx"] = i
                prev["val"] += 1
                counter[c] = prev
        # get the key with greater val
        max_val = -1
        res = []
        for k, data in counter.items():
            if data.get("val") > max_val:
                max_val = data.get("val")
                res = [data]
            elif data.get("val") == max_val:
                res.append(data)
        
        res.sort(key=lambda x: x['idx'])
        res = "".join(list(map(lambda x: s[x['idx']], res)))
        return res

