var mapDigit = map[byte][]rune{
    '2': []rune{'a', 'b', 'c'},
    '3': []rune{'d', 'e', 'f'},
    '4': []rune{'g', 'h', 'i'},
    '5': []rune{'j', 'k', 'l'},
    '6': []rune{'m', 'n', 'o'},
    '7': []rune{'p', 'q', 'r', 's'},
    '8': []rune{'t', 'u', 'v'},
    '9': []rune{'w', 'x', 'y', 'z'},
}

func letterCombinations(digits string) []string {
    if len(digits) == 0 {
        return []string{}
    }

    var result []string

    var backtrack func(index int, currentCombination string)
    
    backtrack = func(i int, currComb string) {
        if i == len(digits) {
            result = append(result, currComb)
            return
        }

        for _, c := range mapDigit[digits[i]] {
            backtrack(i + 1, currComb + string(c))
        }
    }

    backtrack(0, "")

    return result
}
