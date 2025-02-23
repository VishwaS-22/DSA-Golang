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

	// fmt.Print("Enter a target to search in the arr space: ")
	// var t int
	// fmt.Scan(&t)
	// //index := BS_Iterative(t, arr)
	// index := BS_Search(t, arr)
	// if index == -1 {
	// 	fmt.Printf("The target %d is not found.\n", t)
	// } else {
	// 	fmt.Printf("The target %d is found at index %d. \n", t, index)
	// }

	// //Lower Bound
	// fmt.Print("Element to find it's smallest index greater than it or equal to: ")
	// var x int
	// fmt.Scan(&x)
	// index := lower_bound(s, x, arr)
	// fmt.Printf("The smallest index with val >= %d is %d:\n ", x, index)

	// // Upper Bound
	// fmt.Print("Element to find it's smallest index greater than it: ")
	// var x int
	// fmt.Scan(&x)
	// index := upper_bound(s, x, arr)
	// fmt.Printf("The upper bound of %d is %d:\n", x, index)

	// fmt.Print("Enter an element to search it's index: ")
	// var e int
	// fmt.Scan(&e)
	// i := Search_Insertion(s, e, arr)
	// fmt.Printf("The element %d fits in the index: %d\n", e, i)

	fmt.Print("Enter an element to find it's floor and ceil val: ")
	var e int
	fmt.Scan(&e)
	fmt.Printf("The floor val for %d is: %d.\n", e, floor(s, e, arr))
	fmt.Printf("The ceil val for %d is: %d.", e, ceil(s, e, arr))
}
