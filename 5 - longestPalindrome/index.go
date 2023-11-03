package main

import "fmt"

func preProcess(s string) string {
	n := len(s)
	if n == 0 {
		return "^$"
	}
	result := "^"
	for i := 0; i < n; i++ {
		result += "#" + string(s[i])
	}
	result += "#$"
	return result
}

// Manacher's algorithm
func longestPalindrome(s string) string {
	t := preProcess(s)
	n := len(t)
	p := make([]int, n)
	C, R := 0, 0

	for i := 1; i < n-1; i++ {
		iMirror := 2*C - i

		if R > i {
			p[i] = min(R-i, p[iMirror])
		} else {
			p[i] = 0
		}

		for t[i+1+p[i]] == t[i-1-p[i]] {
			p[i]++
		}

		if i+p[i] > R {
			C = i
			R = i + p[i]
		}
	}

	maxLen, centerIndex := 0, 0
	for i := 1; i < n-1; i++ {
		if p[i] > maxLen {
			maxLen = p[i]
			centerIndex = i
		}
	}

	start := (centerIndex - maxLen) / 2
	return s[start : start+maxLen]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	text := "qbmhukucteihghldwdobtvgwwnhflpceiwhbkmvxavmqxedfndegztlpjptpdowwavemasyrjxxnhldnloyizyxgqlhejsdylvkpdzllrzoywfkcamhljminikvwwvqlerdilrdgzifojjlgeayprejhaequyhcohoeonagsmfrqhfzllobwjhxdxzadwxiglvzwiqyzlnamqqsastxlojpcsleohgtcuzzrvwzqugyimaqtorkafyebrgmrfmczwiexdzcokbqymnzigifbqzvfzjcjuugdmvegnvkgbmdowpacyspszvgdapklrhlhcmwkwwqatfswmxyfnxkepdotnvwndjrcclvewomyniaefhhcqkefkyovqxyswqpnysafnydbiuanqphfhfbfovxuezlovokrsocdqrqfzbqhntjafzfjisexcdlnjbhwrvnyabjbshqsxnaxhvtmqlfgdumtpeqzyuvmbkvmmdtywquydontkugdayjqewcgtyajofmbpdmykqobcxgqivmpzmhhcqiyleqitojrrsknhwepoxawpsxcbtsvagybnghqnlpcxlnshihcjdjxxjjwyplnemojhodksckmqdvnzewhzzuswzctnlyvyttuhlreynuternbqonlsfuchxtsyhqlvifcxerzqepthwqrsftaquzuxwcmjjulvjrkatlfqshpyjsbaqwawgpabkkjrtchqmriykbdsxwnkpaktrvmxjtfhwpzmieuqevlodtroiulzgbocrtiuvpywtcxvajkpfmaqckgrcmofkxynjxgvxqvkmhdbvumdaprijkjxxvpqnxakiavuwnifvyfolwlekptxbnctjijppickuqhguvtoqfgepcufbiysfljgmfepwyaxusvnsratn"
	fmt.Println(longestPalindrome(text))
}
