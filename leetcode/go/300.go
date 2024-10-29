//go:build 625
// +build 625

package main

import "fmt"

func lengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, len(nums))
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen
}

func lengthOfLISBS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	tails := make([]int, len(nums))
	size := 0

	for _, num := range nums {
		i, j := 0, size

		for i != j {
			m := (i + j) / 2
			if tails[m] < num {
				i = m + 1
			} else {
				j = m
			}
		}
		tails[i] = num
		if i == size {
			size++
		}
	}
	return size
}

func lengthOfLISBin(nums []int) int {
	res := []int{}

	for _, n := range nums {
		if len(res) == 0 || res[len(res)-1] < n {
			res = append(res, n)
		} else {
			idx := binarySearch(res, n)
			res[idx] = n
		}
	}

	return len(res)
}

func binarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return left
}

func main() {
	testCases := []struct {
		nums     []int
		expected int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{0, 1, 0, 3, 2, 3}, 4},
		{[]int{7, 7, 7, 7, 7, 7, 7}, 1},
		{[]int{}, 0},
		{[]int{1}, 1},
	}

	for i, tc := range testCases {
		result := lengthOfLIS(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed. Expected %d, got %d\n", i+1, tc.expected, result)
		}
	}

	for i, tc := range testCases {
		result := lengthOfLISBS(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed. Expected %d, got %d\n", i+1, tc.expected, result)
		}
	}

	for i, tc := range testCases {
		result := lengthOfLISBin(tc.nums)
		if result == tc.expected {
			fmt.Printf("Test case %d passed\n", i+1)
		} else {
			fmt.Printf("Test case %d failed. Expected %d, got %d\n", i+1, tc.expected, result)
		}
	}
}
