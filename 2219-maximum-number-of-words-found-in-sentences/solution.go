func mostWordsFound(sentences []string) int {
    maxWordsCount := 0
    for _, sentence := range sentences {
        countSingleSpace := 0
        for i, char := range sentence {
            if i == 0 || i == len(sentence) - 1 {
                continue
            } else if char == ' ' {
                // check if previous and next char is non-whitespace
                if sentence[i - 1] != ' ' {
                    countSingleSpace++
                }
            }
        }

        if countSingleSpace > maxWordsCount {
            maxWordsCount = countSingleSpace
        }
    }

    return maxWordsCount + 1
}
