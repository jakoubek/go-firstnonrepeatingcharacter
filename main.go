package main

import "fmt"

func main() {
	str := "ADBCGHIEFKJLADTVDERFSWVGHQWCNOPENSMSJWIERTFB"
	fmt.Println(firstNonRepeatingChar(str))
}

func firstNonRepeatingChar(str string) string {
	m := make(map[rune]uint, len(str))
	for _, r := range str {
		m[r]++
	}

	for _, r := range str {
		if m[r] == 1 {
			return string(r)
		}
	}
	return ""
}
