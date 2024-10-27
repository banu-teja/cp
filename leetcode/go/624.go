//go:build 624
// +build 624

package main

import (
	"fmt"
)

func maxDistance(arrays [][]int) int {
	minVal, maxVal := arrays[0][0], arrays[0][len(arrays[0])-1]
	maxDistance := 0

	for i := 1; i < len(arrays); i++ {
		currMin, currMax := arrays[i][0], arrays[i][len(arrays[i])-1]

		maxDistance = max(maxDistance, abs(currMax-minVal), abs(maxVal-currMin))

		minVal = min(minVal, currMin)
		maxVal = max(maxVal, currMax)
	}

	return maxDistance
}

func max(nums ...int) int {
	if len(nums) == 0 {
		return 0
	}

	maxVal := nums[0]

	for _, num := range nums[1:] {
		if num > maxVal {
			maxVal = num
		}
	}

	return maxVal
}

func min(nums ...int) int {
	minVal := nums[0]

	for _, num := range nums[1:] {
		if num < minVal {
			minVal = num
		}
	}
	return minVal
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func main() {
	testCases := [][][]int{
		{{1, 2, 3}, {4, 5}, {1, 2, 3}},
		{{1}, {1}},
		{{1, 4}, {0, 5}},
	}

	for i, tc := range testCases {
		result := maxDistance(tc)
		fmt.Printf("Test Case %d: %v\nResult: %d\n\n", i+1, tc, result)
	}
}
