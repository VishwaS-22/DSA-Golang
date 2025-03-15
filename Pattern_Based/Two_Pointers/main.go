package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// fmt.Print("Enter size of an arr: ")
	// var s int
	// fmt.Scan(&s)
	// fmt.Print("Enter vals of arr: ")
	// arr := make([]int, s)
	// for i := range arr {
	// 	fmt.Scan(&arr[i])
	// }
	// fmt.Print("Enter target val to look for two sum: ")
	// var t int
	// fmt.Scan(&t)
	// ans := two_sum(arr, t)
	// fmt.Printf("The 1 based index is: %d", ans)
	// fmt.Print("Enter a input string: ")
	// var s string
	// fmt.Scan(&s)
	// if isPalindrome(s) {
	// 	fmt.Println("Yes, It's a Palindrome.")
	// } else {
	// 	fmt.Println("No, it's not a Palindrome.")
	// }
	// fmt.Print("Enter size of an arr: ")
	// var s int
	// fmt.Scan(&s)
	// fmt.Print("Enter vals of arr: ")
	// arr := make([]int, s)
	// for i := range arr {
	// 	fmt.Scan(&arr[i])
	// }
	// ans := threeSum(arr)
	// fmt.Printf("The ans is: %d", ans)
	// fmt.Println("Finding the Max Container that holds more Water: ")
	// var n int
	// fmt.Print("Enter the size of the arr: ")
	// fmt.Scan(&n)
	// fmt.Print("Enter the input which is heights of containers: ")
	// arr := make([]int, n)
	// for i := 0; i < n; i++ {
	// 	fmt.Scan(&arr[i])
	// }
	// sol := maxArea(arr)
	// fmt.Printf("The max length of water that a container can hold is: %d", sol)

	fmt.Println("Reversing the Strings:")
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a sentence: ")
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)
	fmt.Println("The reversed string is: ", reverseWords(str))
}
