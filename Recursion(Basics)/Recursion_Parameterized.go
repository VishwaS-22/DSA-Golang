//Sum of First N nums.

package main

import "fmt"

func Sum_Of_N(n, sum int) {
	if n < 1 {
		fmt.Println(sum)
		return
	}
	Sum_Of_N(n-1, sum+n)
}
