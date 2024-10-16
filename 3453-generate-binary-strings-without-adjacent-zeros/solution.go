func validStrings(n int) []string {
    result := make([]string, 0)
    recurse(n, 0, make([]rune, n), &result)

    return result
}

func recurse(n, i int, result []rune, arr *[]string) {
    if i == n {
        *arr = append(*arr, string(result))
        return
    }

    if i - 1 >= 0 {
        if result[i - 1] == '0' {
            result[i] = '1'
            recurse(n, i + 1, result, arr)
        } else {
            result[i] = '0'
            recurse(n, i+1, result, arr)

            result[i] = '1'
            recurse(n, i+1, result, arr)
        }
    } else {
        result[i] = '0'
        recurse(n, i+1, result, arr)
        result[i] = '1'
        recurse(n, i+1, result, arr)
    }

    return
} 
