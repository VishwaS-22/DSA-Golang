package main

import "fmt"

func N_Nums(i, n int) {
	if i > n {
		return
	}
	fmt.Println(i)
	N_Nums(i+1, n)
}
