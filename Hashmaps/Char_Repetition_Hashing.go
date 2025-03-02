package main

import "fmt"

func Char_Repetition(newString1 string, char1 rune) {
	hasharr := [26]int{}
	for _, v := range newString1 {
		hasharr[v-'a']++
	}
	fmt.Printf("The hashed arr is %v\n", hasharr)

	if char1 >= 'a' && char1 <= 'z' {
		fmt.Printf("The repetition for %c is %v\n", char1, hasharr[char1-'a'])
	} else {
		fmt.Printf("The entered val %c is not in lowercase.\n", char1)
	}
}
