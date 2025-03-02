package main

import (
	"fmt"
)

func main() {
	var arr = []int{99, 88, 77, 66}

	// Optimatal One
	var big = arr[0]

	for _, v := range arr {
		if v > big {
			big = v
		}
	}

	fmt.Printf("The largest val of arr is %v\n", big)

	fmt.Printf("The 2nd largest is %v\n", seclargest(arr))

	fmt.Printf("The 2nd smallest of arr is %v\n", SecondSmallest(arr))

	IsSorted(arr)

	var arr1 = []int{1, 1, 2, 2, 2, 3, 3}
	newArr1 := RemoveDup(arr1)
	fmt.Printf("The duplicate removed array size is %v\n", newArr1)
	fmt.Printf("The dup removed arr is %v\n", arr1[:newArr1])

	arrRotate := []int{1, 2, 3, 4, 5}
	fmt.Printf("The rotated result is %v\n", Left_Rotate(arrRotate))
	//fmt.Println(arrRotate)
	arrR2 := []int{1, 2, 3, 4, 5}
	s := len(arrR2)
	d := 0
	fmt.Printf("The rotated arr by d elements is %v\n", Left_D_Rotate(arrR2, s, d))

	arrZero := []int{1, 0, 2, 0, 3, 4, 5, 0, 6}
	fmt.Printf("The rearranged arr after putting zero at end is %v\n", Zero_To_End(arrZero))

	rec()

	fmt.Println("Printing Name N times using Recursion:")
	var i, n int = 1, 5
	N_Names(i, n)

	fmt.Println("Printing the vals linearly from 1 to N:")
	var a, b = 1, 4
	N_Nums(a, b)

	fmt.Println("Printing from N to One:")
	var a1 = 4
	N_to_One(a1, a1)

	//fmt.Println("Backtracking for N to One:")
	//var a2, b2 = 1, 4
	//Backtracking_One_to_N(a2, b2)

	fmt.Println("Backtracking for One to N:")
	var a3 int = 4
	Backtracking_One_to_N(a3, a3)

	fmt.Println("Sum of N using Parameterized Recursion:")
	a4 := 3
	a5 := 0
	Sum_Of_N(a4, a5)

	fmt.Println("Sum Of N using Functional Recursion:")
	a6 := 3
	fmt.Println(Rec_Func_Sum_Of_N(a6))

	fmt.Println("Factorial of N:")
	a7 := 4
	fmt.Println(Factorial(a7))

	fmt.Println("Reversing an array using recursion based swap:")
	arr6 := []int{1, 2, 3, 4, 5}
	fmt.Printf("Original Arr is %v\n", arr6)
	Array_rev(0, len(arr6), arr6)
	fmt.Printf("Reversed arr is %v\n", arr6)

	fmt.Println("Check if a String is palindrome or not:")
	var input string = "madam"
	fmt.Println(Is_Palindrome(0, len(input), input))

	fmt.Println("Num repetition in arr:")
	arr7 := []int{1, 2, 3, 1, 3}
	s1 := len(arr7)
	n1 := 2
	//fmt.Printf("The sol is %v\n", Num_Repetition(arr7, n1, s1))
	Num_Repetition(arr7, n1, s1)

	// fmt.Println("Pass an arr and find count of repetition:")
	// var size int
	// fmt.Print("Enter the size of the arr: ")
	// fmt.Scan(&size)

	// arr8 := make([]int, size)
	// fmt.Printf("Enter %d elements for the arr:\n", size)
	// for i := 0; i < size; i++ {
	// 	fmt.Printf("Element %d: ", i+1)
	// 	fmt.Scan(&arr8[i])
	// }

	// var inputCheck int
	// fmt.Print("Enter the num of values to check for repetition: ")
	// fmt.Scan(&inputCheck)

	// for i := 0; i < inputCheck; i++ {
	// 	var n int
	// 	fmt.Printf("Enter value %d to check repetition: ", i+1)
	// 	fmt.Scan(&n)
	// 	Num_All_Repetition(arr8, n)
	// }

	// var newString string
	// fmt.Print("Enter an input string: ")
	// fmt.Scan(&newString)

	// var inputChar int
	// fmt.Print("Enter the num of char to check repetition: ")
	// fmt.Scan(&inputChar)

	// for i := 0; i < inputChar; i++ {
	// 	var chars rune
	// 	fmt.Printf("Enter char %d to check repetition: ", i+1)
	// 	fmt.Scanln()
	// 	fmt.Scanf("%c", &chars)
	// 	Char_Repetition(newString, chars)
	// }

	// fmt.Println("Num repetition checking using HashMap.")

	// var in_size int
	// fmt.Print("Enter input arr size: ")
	// fmt.Scan(&in_size)

	// in_nums := make([]int, in_size)

	// fmt.Println("Enter the nums for in_arr:")
	// for i := range in_nums {
	// 	fmt.Scan(&in_nums[i])
	// }

	// // var in_num int
	// // fmt.Println("Enter the num to check: ")
	// // fmt.Scan(&in_num)
	// // Num_rep_hashhmap(in_nums, in_num)
	// var s_num int
	// fmt.Println("Enter the size of num arr to check: ")
	// fmt.Scan(&s_num)
	// in_arr := make([]int, s_num)
	// fmt.Println("Enter the input arr to check: ")
	// for i:=0; i<s_num; i++{
	// 	fmt.Scan(&in_arr[i])
	// }
	// Num_rep_hashmap_arr(in_nums, in_arr)

	//fmt.Println("Sorting an array:")
	fmt.Print("Size of an array is: ")
	var size_arr int
	fmt.Scan(&size_arr)

	size_array := make([]int, size_arr)
	fmt.Print("Get the input values: ")
	for i = 0; i < size_arr; i++ {
		fmt.Scan(&size_array[i])
	}

	//Selection_Sort(size_arr, size_array) // Get min to the Front
	//Bubble_Sort(size_arr, size_array)
	//Insertion_Sort(size_array, size_arr)
	// fmt.Println("The arr before sorting is: ", size_array)
	// // mergeSort(size_array, 0, size_arr-1)
	// // fmt.Print("The sorted arr is: ", size_array)
	// quickSort(size_array, 0, size_arr-1)
	// fmt.Println("The sorted arr is: ", size_array)

	fmt.Println("Enter a target value: ")
	var target int
	fmt.Scan(&target)
	two_sum_pair(size_arr, target, size_array)
}
