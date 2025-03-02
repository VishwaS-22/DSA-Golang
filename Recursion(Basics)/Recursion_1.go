//Recursion example based on a base condition

package main

import (
	"fmt"
)

var c int = 0

func rec() {
	if c == 3 {
		fmt.Println("The base condition met. So ending the recursion.")
		return
	}
	fmt.Printf("Recursive val is %v\n", c)
	c++
	rec()
}
