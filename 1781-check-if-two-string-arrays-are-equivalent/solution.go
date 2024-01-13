func arrayStringsAreEqual(word1 []string, word2 []string) bool {
    word1Combined := ""
    for _, v := range word1 {
        word1Combined += v
    }
    word2Combined := ""
    for _, v := range word2 {
        word2Combined += v
    }

    return word1Combined == word2Combined
}
