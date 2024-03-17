var mapChar = []string{
    ".-",
    "-...",
    "-.-.",
    "-..",
    ".",
    "..-.",
    "--.",
    "....",
    "..",
    ".---",
    "-.-",
    ".-..",
    "--",
    "-.",
    "---",
    ".--.",
    "--.-",
    ".-.",
    "...",
    "-",
    "..-",
    "...-",
    ".--",
    "-..-",
    "-.--",
    "--..",
}

func uniqueMorseRepresentations(words []string) int {
    theMap := map[string]int{}
    
    for _, word := range words {
        transformation := ""
        for _, c := range word {
            transformation += mapChar[c - 'a']
        }

        theMap[transformation]++
    }

    return len(theMap)
}
