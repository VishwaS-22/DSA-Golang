package main

func Is_Palindrome(i, s int, str string) bool {
	if i >= s/2 {
		return true
	}
	if str[i] != str[s-i-1] {
		return false
	}
	return Is_Palindrome(i+1, s, str)
}
