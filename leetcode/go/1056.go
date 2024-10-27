//go:build 1056
// +build 1056

package main

import "fmt"

func confusingNumber(n int) bool {
	rotatedDigits := map[int]int{
		0: 0,
		1: 1,
		6: 9,
		8: 8,
		9: 6,
	}

	rotated := 0
	original := n

	for n > 0 {
		digit := n % 10

		rotatedDigit, isValid := rotatedDigits[digit]
		if !isValid {
			return false
		}

		rotated = rotated*10 + rotatedDigit

		n /= 10
	}
	// fmt.Println(n, rotated)

	return rotated != original
}

func main() {
	testCases := []int{6, 89, 11, 25, 916, 8008}

	for _, tc := range testCases {
		result := confusingNumber(tc)
		fmt.Printf("Is %d a confusing number? %v\n", tc, result)
	}
}
