package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter a string of parantheses to remove it's outermost one: ")
	var str string
	fmt.Scan(&str)
	op := removeOuterParentheses(str)
	fmt.Println("O/P is: ", op)
}
