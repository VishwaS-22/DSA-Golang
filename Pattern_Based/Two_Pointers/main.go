package main

import "fmt"

func main() {
	fmt.Print("Enter size of an arr: ")
	var s int
	fmt.Scan(&s)
	fmt.Print("Enter vals of arr: ")
	arr := make([]int, s)
	for i := range arr {
		fmt.Scan(&arr[i])
	}
	fmt.Print("Enter target val to look for two sum: ")
	var t int
	fmt.Scan(&t)
	ans := two_sum(arr, t)
	fmt.Printf("The 1 based index is: %d", ans)
}
