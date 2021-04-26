package src

import (
	"fmt"
)

func isPalindrome(x int) bool {
	// negative numbers are not palindrome integer.
	if x < 0 {
		return false
	}
	return isPalindromeString(fmt.Sprintf("%d", x))
}

func isPalindromeString(x string) bool {
	if len(x) == 0 {
		return true
	}
	if len(x) == 1 {
		return true
	}
	if x[0] != x[len(x)-1] {
		return false
	}
	return isPalindromeString(x[1 : len(x)-1])
}
