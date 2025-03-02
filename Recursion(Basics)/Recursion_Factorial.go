package main

//import "fmt"

func Factorial(n int) int {
	if n == 0 {
		return 1 // No zero coz n * 0 will be zero
	}
	return n * Factorial(n-1)
}
