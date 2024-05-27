func lengthOfLongestSubstring(s string) int {
    max := 0
    left := 0
    right := 0
    set := map[byte]int{}

    for right = 0; right < len(s); right++ {
        set[s[right]]++

        for set[s[right]] > 1 {
            // reset start (left) index
            set[s[left]]--
            left++
        }

        if right - left + 1 > max {
            max = right - left + 1
        }
    }

    return max
}
