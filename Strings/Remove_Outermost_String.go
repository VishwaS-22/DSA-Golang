package main

import "strings"

func removeOuterParentheses(S string) string {
	var result strings.Builder // Using strings.Builder for efficiency
	balance := 0               // Balance counter

	for _, ch := range S {
		// If we encounter a '('
		if ch == '(' {
			// If balance is non-zero, this '(' is not an outer parenthesis.
			if balance > 0 {
				result.WriteRune(ch)
			}
			balance++ // Increase the balance
		} else { // ch == ')'
			balance-- // Decrease the balance
			// If balance is non-zero after decreasing, then this ')' is not an outer parenthesis.
			if balance > 0 {
				result.WriteRune(ch)
			}
		}
	}
	return result.String()
}
