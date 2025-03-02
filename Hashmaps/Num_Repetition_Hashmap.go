package main

import "fmt"

func Num_rep_hashhmap(arr []int, val int) {

	hashmap := make(map[int]int)
	fmt.Printf("The hashmap default val before computing %v.\n", hashmap)

	for _, v := range arr {
		hashmap[v]++
	}

	fmt.Printf("The val after hashed is %v.\n", hashmap)

	if c, ok := hashmap[val]; ok {
		fmt.Printf("The val %d is repeated %v times.\n", val, c)
	} else {
		fmt.Printf("The num %d does not exists.", val)
	}
}

func Num_rep_hashmap_arr(arr1 []int, c_arr []int){
	hashmap := make(map[int]int)
	
	// Insert numbers and compute frequency
	for _, num := range arr1 {
		hashmap[num]++
	}

	// Process queries and print results
	for _, q := range c_arr {
		fmt.Printf("Frequency of %d: %d\n", q, hashmap[q])
	}
}