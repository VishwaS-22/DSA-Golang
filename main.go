package main

import "fmt"

func main() {
	fmt.Print("Enter the size of the array: ")
	var s int
	fmt.Scan(&s)
	fmt.Print("Enter the values for an array: ")
	arr := make([]int, s)
	for v := range arr {
		fmt.Scan(&arr[v])
	}
	fmt.Print("Enter a target to search in the arr space: ")
	var t int
	fmt.Scan(&t)
	//index := BS_Iterative(t, arr)
	index := BS_Search(t, arr)
	if index == -1 {
		fmt.Printf("The target %d is not found.\n", t)
	} else {
		fmt.Printf("The target %d is found at index %d. \n", t, index)
	}
	
}
