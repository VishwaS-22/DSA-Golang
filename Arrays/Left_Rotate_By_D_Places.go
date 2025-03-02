//Left rotate the element in an arr by D places

package main

import "fmt"

func Left_D_Rotate(arr []int, n, d int) []int {
	d = d % n

	//if d == 0 {
	// 	fmt.Println("No rotation needed, returning original array.")
	// 	return arr
	// }
	
	t := make([]int, d)
	//n := len(arr)
	for i := range arr[:d] {
		t[i] = arr[i]
	}
	fmt.Printf("Temp is %v\n", t)
	// copy(t,arr[:d])

	for i := d; i < n; i++ {
		arr[i-d] = arr[i]
	}

	for i := n - d; i < n; i++ {
		arr[i] = t[i-(n-d)]
	}

	return arr
}
