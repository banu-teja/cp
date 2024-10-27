//go:build 1427
// +build 1427

package main

import "fmt"

func stringShift(s string, shift [][]int) string {
	netShift := 0
	for _, op := range shift {
		if op[0] == 0 {
			netShift -= op[1]
		} else {
			netShift += op[1]
		}
	}

	netShift %= len(s)

	if netShift < 0 {
		netShift += len(s)
	}

	return s[len(s)-netShift:] + s[:len(s)-netShift]
}

func main() {
	testCases := []struct {
		s     string
		shift [][]int
		want  string
	}{
		{"abc", [][]int{{0, 1}, {1, 2}}, "cab"},
		{"abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}}, "efgabcd"},
		{"abcdefg", [][]int{{1, 1}, {1, 1}, {0, 2}, {1, 3}, {0, 7}}, "efgabcd"},
	}

	for i, tc := range testCases {
		got := stringShift(tc.s, tc.shift)
		if got != tc.want {
			fmt.Printf("Test case %d failed. Got: %s, Want: %s\n", i+1, got, tc.want)
		} else {
			fmt.Printf("Test case %d passed.\n", i+1)
		}
	}
}
