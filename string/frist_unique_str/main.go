package main

import (
	"fmt"
)

func main() {
	input := "leetcode"

	fmt.Println(firstUniqChar(input))
}

func firstUniqChar(s string) int {
	if len(s) == 1 {
		return 0
	}

	data := map[rune]int{}
	res := -1

	for _, str := range s {
		if _, ok := data[str]; ok {
			data[str]++
			continue
		}
		data[str] = 1
	}

	for v, str := range s {
		if data[str] == 1 {
			res = v
			break
		}
	}

	return res
}

func firstUniqChar2(s string) int {
	var cahrlist [26]int

	for _, c := range s {
		cahrlist[c-'a']++
	}

	for i, c := range s {
		if cahrlist[c-'a'] == 1 {
			return i
		}
	}

	return -1
}
