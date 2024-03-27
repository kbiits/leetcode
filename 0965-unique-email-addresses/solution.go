func numUniqueEmails(emails []string) int {
    mapEmail := map[string]int{}

    for _, email := range emails {

        var localPart []byte

        var j int
		// find all local part
        for email[j] != '@' && email[j] != '+' {
			if email[j] != '.' {
				localPart = append(localPart, email[j])
			}
			j++
		}

        // find @ index and ignore after + sign
		for email[j] != '@' {
			j++
		}
		
        // concat local and domain part
		addr := string(localPart) + string(email[j:])
		mapEmail[addr]++
    }

    return len(mapEmail)
}

