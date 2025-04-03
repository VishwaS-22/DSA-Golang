package main 

func isHappy(n int) bool {
    slow, fast := n, n

    for {
        slow = sumOfSquares(slow)         // Move slow one step
        fast = sumOfSquares(sumOfSquares(fast)) // Move fast two steps

        if slow == 1 || fast == 1 {
            return true // Reached 1, it's a happy number
        }

        if slow == fast {
            return false // Cycle detected, not a happy number
        }
    }
}
func sumOfSquares(n int) int {
    sum := 0
    for n > 0 {
        digit := n % 10
        sum += digit * digit
        n /= 10
    }
    return sum
}