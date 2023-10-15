package main

import (
	"fmt"
	"strings"
)

func main() {
	haystack := "leetcode"
	needle := "tc"
	fmt.Println(strStr3(haystack, needle))
	fmt.Println(strStr(haystack, needle))
}

func strStr(haystack string, needle string) int {
	fund_index := 0
	res := -1
	for i := 0; i < len(haystack); i++ {
		str_i := 0
		for j := 0; j < len(needle); j++ {
			if i+j >= len(haystack) {
				continue
			}
			if haystack[i+j] == needle[j] {
				str_i++
			}
		}
		if str_i == len(needle) {
			fund_index = i
			res = 0
			break
		}
	}

	if fund_index != 0 && res == 0 {
		res = fund_index
	}

	return res
}

// 別人的寫法
func strStr2(haystack string, needle string) int {
	for i := 0; i <= len(haystack)-len(needle); i++ {
		k := 0
		for j := 0; j < len(needle); j++ {
			if haystack[i+j] != needle[j] {
				break
			}
			k++
		}
		if k == len(needle) {
			return i
		}
	}
	return -1
}

// 哀哀...原來還是有函示
func strStr3(haystack string, needle string) int {
	return strings.Index(haystack, needle)
}
