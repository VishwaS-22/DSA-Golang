package main

import "fmt"

func Num_Repetition(arr []int, n, size int) {
	hash := [10]int{}
	fmt.Printf("The default hash arr is %v\n", hash)
	// for i:=0; i< size; i--{
	// 	hash[arr[i]] +=1
	// }

	for _, v := range arr {
		hash[v] += 1
	}

	fmt.Printf("The hash arr after pre-computation is %v\n", hash)

	if n >= 0 {
		fmt.Printf("The repeated count is %v\n", hash[n])
	}
}

func Num_All_Repetition(arr1 []int, n1 int) {
	hash := [10]int{}
	fmt.Printf("The default val of hash arr is %v\n", hash)

	for _, v := range arr1 {
		hash[v] += 1
	}

	fmt.Printf("The hash arr after pre-computation is %v\n", hash)

	if n1 >= 0 && n1 < len(hash) {
		fmt.Printf("The repeated count for %d is %v\n", n1, hash[n1])
	} else {
		fmt.Printf("Number %d is out of range for hash computation.\n", n1)
	}
}
