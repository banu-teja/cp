// go:build 1574
//go:build 1574
// +build 1574

package main

import "fmt"

func findLengthOfShortestSubarray(arr []int) int {
	n := len(arr)
	left := 0
	for left < n-1 && arr[left] <= arr[left+1] {
		left++
	}

	if left == n-1 {
		return 0
	}

	right := n - 1
	for right > left && arr[right-1] <= arr[right] {
		right--
	}

	result := min(n-left-1, right)

	i, j := 0, right
	for i <= left && j < n {
		if arr[i] <= arr[j] {
			result = min(result, j-i-1)
			i++
		} else {
			j++
		}
	}

	return result
}

func main() {
	testCases := []struct {
		arr      []int
		expected int
	}{
		{[]int{1, 2, 3, 10, 4, 2, 3, 5}, 3},
		{[]int{5, 4, 3, 2, 1}, 4},
		{[]int{1, 2, 3}, 0},
		{[]int{1}, 0},
	}

	for i, tc := range testCases {
		result := findLengthOfShortestSubarray(tc.arr)
		fmt.Printf("Test case %d: Expected %d, Got %d\n", i+1, tc.expected, result)
		if result != tc.expected {
			fmt.Println("FAILED")
		} else {
			fmt.Println("PASSED")
		}
	}
}
