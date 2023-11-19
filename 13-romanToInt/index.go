package main

func romanToInt(s string) int {
	romanMap := make(map[byte]int)

	romanMap['I'] = 1
	romanMap['V'] = 5
	romanMap['X'] = 10
	romanMap['L'] = 50
	romanMap['C'] = 100
	romanMap['D'] = 500
	romanMap['M'] = 1000

	num := 0

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && romanMap[s[i+1]] > romanMap[s[i]] {
			num -= romanMap[s[i]]
			continue
		}
		num += romanMap[s[i]]
	}

	return num
}

func main() {
	println(romanToInt("IV"))
}
