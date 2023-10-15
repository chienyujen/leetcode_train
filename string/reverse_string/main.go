package main

func main() {
	input := []byte{'H', 'a', 'p', 'p', 'g', 'Y'}
	reverseString2(input)
}

func reverseString(s []byte) {
	n := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		temp := s[i]
		s[i] = s[n-i]
		s[n-i] = temp
	}
}

func reverseString2(s []byte) {
	i := 0
	n := len(s) - 1
	for i < n {
		s[i], s[n] = s[n], s[i]
		i++
		n--
	}
}
