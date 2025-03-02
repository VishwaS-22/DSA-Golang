package main

// func main() {
// 	arr := []int{10, 6, 8, 4, 0, 2}
// 	seclargest(arr)

// }

// Optimal Sol
func seclargest(arr1 []int) int {
	var lar = arr1[0]
	seclar := -1
	for _, v := range arr1[1:] { //Starts from 2nd element
		if v > lar {
			seclar = lar
			lar = v
		} else if v < lar && v > seclar {
			seclar = v
		}
	}
	return seclar
}
