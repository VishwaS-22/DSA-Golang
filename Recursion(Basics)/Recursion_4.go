package main

import "fmt"

func N_to_One(i, n int) {
	if i < 1 {
		return
	}
	fmt.Println(i)
	N_to_One(i-1, n)
}
