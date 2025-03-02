// Max points you can obtain from cards,
// Only take consecutive K vals from starting, form reverse of arr.
package main

func Max_Points(arr []int, k int) int {
	leftSum := 0
	rightSum := 0
	maxSum := 0

	for i := 0; i <= k-1; i++ {
		leftSum += arr[i]
		maxSum = leftSum
	}
	rightIndex := len(arr) - 1
	for i := k - 1; i >= 0; i-- {
		leftSum = leftSum - arr[i]
		rightSum = rightSum + arr[rightIndex]
		rightIndex = rightIndex - 1
		newMaxSum := rightSum + leftSum
		if newMaxSum > maxSum {
			maxSum = newMaxSum
		}
	}
	return maxSum
}