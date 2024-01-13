func countMatches(items [][]string, ruleKey string, ruleValue string) int {
    count := 0
    ruleIdx := 0
    if ruleKey == "color" {
        ruleIdx = 1
    } else if ruleKey == "name" {
        ruleIdx = 2
    }

    for _, item := range items {
        if item[ruleIdx] == ruleValue {
            count++
        }
    }

    return count
}
