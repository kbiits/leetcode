func countAndSay(n int) string {
    if n == 1 {
        return "1"
    }

    str := countAndSay(n - 1)
    // fmt.Printf("n = %d, str = %s\n", n, str)
    return buildRLE(str)
}


func buildRLE(str string) string {
	if len(str) == 0 {
		return ""
	}

	if len(str) == 1 {
		return "1" + str
	}

	c := str[0]
	newStr := ""
	count := 0
	for i := 0; i < len(str); i++ {
		if i < len(str)-1 && str[i] != str[i+1] {
			count++
			newStr += fmt.Sprintf("%d%c", count, c)
			count = 0
			c = str[i+1]
		} else {
			count++
		}

		if i == len(str)-1 {
			newStr += fmt.Sprintf("%d%c", count, c)
		}
	}

	return newStr
}

