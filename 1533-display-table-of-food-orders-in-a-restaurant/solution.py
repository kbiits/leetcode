class Solution:
    def displayTable(self, orders: List[List[str]]) -> List[List[str]]:
        header = ["Table"]
        map_table = {}
        item_set = set()

        for order in orders:
            item = order[2]
            
            item_set.add(item)

            table_no = order[1]
            if not (table_no in map_table):
                map_table[table_no] = {}

            if item in map_table[table_no]:
                map_table[table_no][item] += 1
            else:
                map_table[table_no][item] = 1
    
        result = [header]
        sorted_item = sorted(item_set)

        for item in sorted_item:
            header.append(item)

        for k, v in sorted(map_table.items(), key=lambda item: int(item[0])):
            row = [k]
            for item in sorted_item:
                row.append(str(v[item] if item in v else "0"))
            result.append(row)

        return result
