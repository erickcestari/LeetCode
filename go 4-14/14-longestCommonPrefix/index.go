package main

import (
	"fmt"
)

func minString(a, b string) string {
	if a < b {
		return a
	}
	return b
}

func minArrayString(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	if len(strs) == 2 {
		return minString(strs[0], strs[1])
	}

	min := strs[0]
	isOdd := len(strs) % 2
	for i := 1; i < (len(strs)+isOdd)/2; i++ {
		left := strs[i]
		right := strs[len(strs)-i]
		min = minString(min, minString(left, right))
	}

	return min
}

func longestCommonPrefix(strs []string) string {
	longestPrefix := ""
	shortestString := minArrayString(strs)

	for i := 0; i < len(shortestString); i++ {
		for j := 0; j < len(strs); j++ {
			if shortestString[i] != strs[j][i] {
				return longestPrefix
			}
		}
		longestPrefix += string(shortestString[i])
	}

	return longestPrefix
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"abab", "aba", ""}))
}
