package main

func intToRoman(num int) string {
	return getRoman(num, "")
}

func getRoman(num int, roman string) string {
	if num == 0 {
		return roman
	}

	if num-5000 >= 0 {
		return getRoman(num-5000, roman+"V̅")
	}

	if num-4000 >= 0 {
		return getRoman(num-4000, roman+"MV̅")
	}

	if num-1000 >= 0 {
		return getRoman(num-1000, roman+"M")
	}

	if num-900 >= 0 {
		return getRoman(num-900, roman+"CM")
	}

	if num-500 >= 0 {
		return getRoman(num-500, roman+"D")
	}

	if num-400 >= 0 {
		return getRoman(num-400, roman+"CD")
	}

	if num-100 >= 0 {
		return getRoman(num-100, roman+"C")
	}

	if num-90 >= 0 {
		return getRoman(num-90, roman+"XC")
	}

	if num-50 >= 0 {
		return getRoman(num-50, roman+"L")
	}

	if num-40 >= 0 {
		return getRoman(num-40, roman+"XL")
	}

	if num-10 >= 0 {
		return getRoman(num-10, roman+"X")
	}

	if num-9 >= 0 {
		return getRoman(num-9, roman+"IX")
	}

	if num-5 >= 0 {
		return getRoman(num-5, roman+"V")
	}

	if num-4 >= 0 {
		return getRoman(num-4, roman+"IV")
	}

	if num-1 >= 0 {
		return getRoman(num-1, roman+"I")
	}

	return roman
}

func main() {
	println(intToRoman(8123))
}
