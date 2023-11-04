package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func reverseString(input string) string {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverse(input int) int {
	strInput := strconv.Itoa(input)

	reversedStr := reverseString(strInput)

	if strings.HasSuffix(reversedStr, "-") {
		reversedStr = strings.TrimSuffix(reversedStr, "-")
		reversedStr = "-" + reversedStr
	}

	reversedInt, err := strconv.Atoi(reversedStr)
	if err != nil {
		fmt.Println("Error:", err)
		return 0
	}

	maxInt32 := math.MaxInt32

	if reversedInt > maxInt32 || reversedInt < -maxInt32 {
		return 0
	}

	return reversedInt
}

func main() {

	fmt.Println(reverse(-2147483648))
}
