package main

func isAlphaNumeric(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9')
}

func toLower(ch byte) byte {
	if ch >= 'A' && ch <= 'Z' {
		return ch - 'A' + 'a'
	}
	return ch
}

func isPalindrome(s string) bool {
	l, r := 0, len(s)-1

	for l < r {
		for l < r && !isAlphaNumeric(s[l]) {
			l++
		}
		for l < r && !isAlphaNumeric(s[r]) {
			r--
		}

		if toLower(s[l]) != toLower(s[r]) {
			return false
		}
		l++
		r--
	}
	return true
}
