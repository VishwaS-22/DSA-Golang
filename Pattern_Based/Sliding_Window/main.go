package main

import "fmt"

func main() {
	fmt.Print("Enter an size for an array: ")
	var s int
	fmt.Scan(&s)

	fmt.Print("Get the input vals for array: ")
	arr := make([]int, s)
	for i := range arr {
		fmt.Scan(&arr[i])
	}

	fmt.Print("Enter value for k length: ")
	var k int
	fmt.Scan(&k)

	// sol := maxSubarrsum(arr, k)
	// fmt.Print("The ans is: ", sol)

	sol1 := maxSubarrCondition(arr, k)
	fmt.Print("The ans is: ", sol1)
}
