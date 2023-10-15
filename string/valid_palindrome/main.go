package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	input := "0P"
	fmt.Println(isPalindrome2(input))
}

func isPalindrome(s string) bool {
	low_str := strings.ToLower(s)
	tempstr := ""
	tempstr2 := ""
	for _, c := range low_str {
		if (c >= '0' && c <= '9') || (c >= 'a' && c <= 'z') {
			tempstr = fmt.Sprintf("%v%v", tempstr, string(c))
		}
		continue
	}
	if len(tempstr) == 0 {
		return true
	}
	for i := len(tempstr) - 1; i >= 0; i-- {
		tempstr2 = fmt.Sprintf("%v%v", tempstr2, string(tempstr[i]))
	}
	if tempstr == tempstr2 {
		return true
	}
	return false
}

func isPalindrome2(s string) bool {
	low_str := strings.Map(delOther, s)
	fmt.Println(low_str)
	i := 0
	j := len(low_str) - 1

	for i < j {
		if low_str[i] == low_str[j] {
			i++
			j--
		} else {
			return false
		}
	}

	return true
}

func delOther(r rune) rune {
	if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
		return -1
	}

	return unicode.ToLower(r)
}

func isPalindrome3(s string) bool {
	i := 0
	j := len(s) - 1
	for i < j {
		s_i := strings.ToLower(string(s[i]))
		s_j := strings.ToLower(string(s[j]))

		if !isdelchat(rune(s[i])) {
			i++
			continue
		}
		if !isdelchat(rune(s[j])) {
			j--
			continue
		}

		if s_i == s_j {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func isdelchat(x rune) bool {
	return unicode.IsLetter(x) || unicode.IsNumber(x)
}
