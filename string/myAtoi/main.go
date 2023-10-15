package main

import (
	"fmt"
	"math"
	"strconv"
	"unicode"
)

func main() {
	input := "-9999999999999"

	fmt.Println(myAtoi2(input))
}

func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}

	temp := ""
	fristword := 0
	for i := 0; i < len(s); i++ {
		if fristword != 0 && len(temp) != 0 && !unicode.IsNumber(rune(s[i])) {
			_, err := strconv.Atoi(temp)
			if (s[i] == '-' || s[i] == '+') && err != nil {
				return 0
			}
			break
		}

		if s[i] == ' ' && fristword == 0 {
			continue
		}

		if (s[i] == '-' || s[i] == '+') && fristword == 0 {
			fristword++
			temp = fmt.Sprintf("%v%v", temp, string(s[i]))
		}

		if !unicode.IsNumber(rune(s[i])) && fristword == 0 {
			return 0
		}

		if unicode.IsNumber(rune(s[i])) {
			temp = fmt.Sprintf("%v%v", temp, string(s[i]))
			fristword++
		}
	}

	if temp == "" {
		return 0
	}

	res, _ := strconv.Atoi(temp)
	if res > (1<<31)-1 {
		return (1 << 31) - 1
	}
	if res < -(1 << 31) {
		return -(1 << 31)
	}
	return res
}

func myAtoi2(s string) int {
	answer := 0
	sign := 1
	isNumber := false
	for _, c := range s {
		if c == ' ' && !isNumber {
			continue
		}

		if c == '-' && sign == 1 && !isNumber {
			isNumber = true
			sign = -1
			continue
		} else if c == '+' && sign == 1 && !isNumber {
			isNumber = true
			sign = 1
			continue
		} else if !(c >= '0' && c <= '9') {
			break
		}

		isNumber = true
		answer = answer*10 + sign*int(c-'0')
		//if sign == -1 && answer > 0 {
		//	answer *= -1
		//}

		if answer > math.MaxInt32 {
			return math.MaxInt32
		}
		if answer < math.MinInt32 {
			return math.MinInt32
		}
	}

	return answer
}
