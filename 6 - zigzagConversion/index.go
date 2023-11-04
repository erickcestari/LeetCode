package main

import "fmt"

func convert(s string, numRows int) string {
	topJump := (numRows-2)*2 + 2
	if topJump <= 0 {
		return s
	}

	size := len(s)
	resultArray := make([]byte, size)
	j := 0
	// First row
	for i := 0; i < size; i += topJump {
		resultArray[j] = s[i]
		j++
	}
	// Middle rows
	for row := 1; row < numRows-1; row++ {
		for i := 0; i-row < size; i += topJump {
			iLeft := i - row
			iRight := i + row
			if iLeft > 0 {
				resultArray[j] = s[iLeft]
				j++
			}
			if iRight < size {
				resultArray[j] = s[iRight]
				j++
			}
		}
	}
	// Last row
	for i := topJump - numRows + 1; i < size; i += topJump {
		resultArray[j] = s[i]
		j++
	}
	return string(resultArray)
}

func main() {
	txt := "PAYPALISHIRING"
	fmt.Println(convert(txt, 3))
}

// PAYPALISHIRING

//P   A   H   N
//A P L S I I G
//Y   I   R
