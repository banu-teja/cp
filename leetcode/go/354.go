//go:build 354
// +build 354

package main

import (
	"fmt"
	"sort"
)

// failing on last 2 testcases
func maxEnvelopes(envelopes [][]int) int {
	sort.Slice(envelopes, func(i, j int) bool {
		return envelopes[i][0] < envelopes[j][0]
	})

	dp := make([]int, len(envelopes))

	maxEnv := 1

	for i := 0; i < len(envelopes); i++ {
		dp[i] = 1

		for j := 0; j < i; j++ {
			if envelopes[j][0] < envelopes[i][0] && envelopes[j][1] < envelopes[i][1] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}

		maxEnv = max(dp[i], maxEnv)
	}

	return maxEnv

}

func main() {
	tests := []struct {
		envelopes [][]int
		expected  int
	}{
		{
			envelopes: [][]int{
				{5, 4},
				{6, 4},
				{6, 7},
				{2, 3},
			},
			expected: 3,
		},
		{
			envelopes: [][]int{
				{1, 1},
				{1, 1},
				{1, 1},
			},
			expected: 1,
		},
		{
			envelopes: [][]int{
				{1, 2},
				{2, 3},
				{3, 4},
			},
			expected: 3,
		},
	}

	for _, test := range tests {
		actual := maxEnvelopes(test.envelopes)
		fmt.Printf("expected %d, got %d\n", test.expected, actual)
	}
}
