package main

import (
	"fmt"
	"strconv"
)

func main() {
	input := 153423646

	fmt.Println(reverse2(input))
}

func reverse(x int) int {
	if x > (1<<31) || x < -(1<<31) {
		return 0
	}
	negat := false
	if x < 0 {
		negat = true
		x = x * -1
	}
	str := strconv.Itoa(x)
	bytes := []byte(str)
	i := 0
	j := len(str) - 1
	for i < j {
		bytes[i], bytes[j] = bytes[j], bytes[i]
		i++
		j--
	}
	str = string(bytes)
	if negat {
		str = fmt.Sprintf("-%s", str)
	}
	res, _ := strconv.Atoi(str)
	if res > (1<<31) || res < -(1<<31) {
		return 0
	}
	return res
}

func reverse2(x int) int {
	var negatif bool
	if x < 0 {
		negatif = true
		x = x * -1
	}
	res := 0
	for x > 0 {
		rest := x % 10
		res = res*10 + rest
		x = (x - rest) / 10
		if res > 2147483647 {
			return 0
		}
	}
	if negatif {
		res = res * -1
	}

	return res
}
