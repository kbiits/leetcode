var (
    mapPairs = map[rune]rune{
        ')': '(',
        '}': '{',
        ']': '[',
    }
)

func isValid(s string) bool {
    stack := []rune{}
    for _, v := range s {
        switch v {
            case ')', '}', ']':
                if len(stack) - 1 < 0 {
                    return false
                }

                popped := stack[len(stack) - 1]
                if mapPairs[v] != popped {
                    return false
                }
                stack = stack[:len(stack) - 1]
                continue
        }

        stack = append(stack, v)
    }

    return len(stack) == 0
}
