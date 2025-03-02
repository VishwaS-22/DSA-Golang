package main

import "fmt"

func N_Names(i, n int) {
	if i > n {
		return
	}
	fmt.Println("Vishwa")
	N_Names(i+1, n)
}
