//go:build 368
// +build 368

package main

import (
	"fmt"
	"sort"
)

// largestDivisibleSubset returns the largest subset of nums such that every pair of elements in the subset is divisible.
func largestDivisibleSubset(nums []int) []int {
	if len(nums) == 0 {
		return []int{}
	}

	sort.Ints(nums)
	dp := make([][]int, len(nums))
	for i := range dp {
		dp[i] = []int{nums[i]}
	}

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i]%nums[j] == 0 && len(dp[j])+1 > len(dp[i]) {
				dp[i] = append([]int{}, dp[j]...)
				dp[i] = append(dp[i], nums[i])
			}
		}
	}
	maxLen, maxIndex := 0, 0
	for i, subset := range dp {
		if len(subset) > maxLen {
			maxLen = len(subset)
			maxIndex = i
		}
	}

	return dp[maxIndex]
}

func main() {
	testCases := [][]int{
		{1, 2, 3},
		{1, 2, 4, 8},
		{3, 4, 16, 8},
		{2, 3, 4, 9, 8},
	}

	for i, tc := range testCases {
		result := largestDivisibleSubset(tc)
		fmt.Printf("Test Case %d: %v\nResult: %v\n\n", i+1, tc, result)
	}
}
