package main

import "fmt"

func Backtracking_One_to_N(i, n int) {
	if i < 1 {
		return
	}
	Backtracking_One_to_N(i-1, n)
	fmt.Println(i)
}
