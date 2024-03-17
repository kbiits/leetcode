import (
    "slices"
)

func groupAnagrams(strs []string) [][]string {
    groupMap := map[string][]string{}

    for _, v := range strs {
        vRune := []rune(v)
        slices.Sort(vRune)

        groupMap[string(vRune)] = append(groupMap[string(vRune)], v)
    }

    groups := [][]string{}
    for _, v := range groupMap {
        groups = append(groups, v)
    }

    return groups
}

