//go:build 625
// +build 625

package main

import "fmt"

func wiggleSort(nums []int) {
	for i := 1; i < len(nums); i++ {
		if (i%2 == 1 && nums[i] < nums[i-1]) || (i%2 == 0 && nums[i] > nums[i-1]) {
			nums[i], nums[i-1] = nums[i-1], nums[i]
		}
	}
}

func main() {
	testCases := [][]int{
		{3, 5, 2, 1, 6, 4},
		{6, 6, 5, 6, 3, 8},
		{1, 5, 1, 1, 6, 4},
	}

	for _, nums := range testCases {
		fmt.Printf("Before: %v\n", nums)
		wiggleSort(nums)
		fmt.Printf("After:  %v\n\n", nums)
	}
}
