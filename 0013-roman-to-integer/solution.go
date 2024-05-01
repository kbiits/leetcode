var mapRoman = map[rune]int{
    'I': 1,
    'V': 5,
    'X': 10,
    'L': 50,
    'C': 100,
    'D': 500,
    'M': 1000,
}

func romanToInt(s string) int {
    
    sum := mapRoman[rune(s[0])]
    for i := 1; i < len(s); i++ {
        curr := rune(s[i])
        sum += mapRoman[curr]
        if mapRoman[curr] > mapRoman[rune(s[i - 1])] {
            sum -= 2 * mapRoman[rune(s[i - 1])]
        }
    }

    return sum
}
