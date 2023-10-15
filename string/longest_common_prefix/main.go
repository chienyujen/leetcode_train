package main

import "fmt"

func main() {
	input := []string{"c", "acc", "ccc"}
	fmt.Println(longestCommonPrefix(input))
}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 1 {
		return strs[0]
	}
	temp := ""
	for i := 1; i < len(strs); i++ {
		hay := strs[0]
		if temp != "" {
			hay = temp
			temp = ""
		}

		for k := 0; k < len(hay); k++ {
			if k >= len(strs[i]) {
				continue
			}
			if hay[k] != strs[i][k] {
				break
			}
			temp = fmt.Sprintf("%v%v", temp, string(hay[k]))
		}

		if temp == "" {
			return ""
		}

	}
	//fmt.Println(temp)
	return temp
}

func longestCommonPrefix2(strs []string) string {
	if len(strs[0]) == 0 {
		return ""
	}

	res := ""
	indx := 0

	for i := 0; i < len(strs[0]); i++ {
		for j := 1; j < len(strs); j++ {
			if indx >= len(strs[j]) || indx >= len(strs[j-1]) || strs[j][i] != strs[j-1][i] {
				return res
			}
		}

		res += string(strs[0][i])
		indx += 1
	}
	return res
}
