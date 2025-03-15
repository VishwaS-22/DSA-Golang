package main

import (
	//"fmt"
	"strings"
)

// func reverseWords(s string) string {
// 	words := strings.Fields(s)
// 	reverse(words)
// 	return strings.Join(words, " ")
// }

// func reverse(w []string) {
// 	l, r := 0, len(w)-1
// 	for l < r {
// 		w[l], w[r] = w[r], w[l]
// 		l++
// 		r--
// 	}
// }

//In-Place Approach (O(1) Extra Space)

func reverseWords(s string) string {
	str := []byte(strings.TrimSpace(s))
	// fmt.Println("Byte Converted String after removing Leading/Trailing spaces is: ", str)
	// fmt.Println("Byte Converted String as String is: ", string(str))
	reverse(str, 0, len(str)-1)
	// fmt.Println("Reversed Converted Byte is: ", str)
	// fmt.Println("Reversed Converted Byte as String is: ", string(str))

	start := 0
	for end := 0; end <= len(str); end++ {
		if end == len(str) || str[end] == ' ' {
			reverse(str, start, end-1)
			start = end + 1
		}
	}
	return cleanSpaces(str)
}

func reverse(str []byte, l, r int) {
	for l < r {
		str[l], str[r] = str[r], str[l]
		l++
		r--
	}
}

func cleanSpaces(str []byte) string {
	res := []byte{}
	spaceFound := false

	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			if !spaceFound {
				res = append(res, ' ')
				spaceFound = true
			}
		} else {
			res = append(res, str[i])
			spaceFound = false
		}
	}
	return string(res)
}
