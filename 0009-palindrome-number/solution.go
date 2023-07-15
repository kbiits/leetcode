func isPalindrome(x int) (out bool) {
	if x == 0 {
		return true
	}
	if x < 0 || x%10 == 0 {
		return false
	}

	sumHalfFromRight := 0
	for x > sumHalfFromRight {
		lastDigit := x % 10
		sumHalfFromRight = sumHalfFromRight*10 + lastDigit

		// reduce x
		x = x / 10
	}

	return x == sumHalfFromRight || x == (sumHalfFromRight / 10)
}

