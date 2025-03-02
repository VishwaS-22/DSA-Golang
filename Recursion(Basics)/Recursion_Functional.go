//Sum of First N nums.

package main

func Rec_Func_Sum_Of_N(n int) int {
	if n < 1 {
		return 0
	}
	return n + Rec_Func_Sum_Of_N(n-1)
}
