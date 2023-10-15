package main

import "fmt"

func main() {
	input1 := "anagram"
	input2 := "nagaram"
	fmt.Println(isAnagram(input1, input2))
}

func isAnagram(s string, t string) bool {
	input1 := [26]int{}
	input2 := [26]int{}
	for _, c := range s {
		input1[c-'a']++
	}
	for _, c := range t {
		input2[c-'a']++
	}

	for v, c := range input1 {
		if input2[v] != c {
			return false
		}
	}
	return true
}
