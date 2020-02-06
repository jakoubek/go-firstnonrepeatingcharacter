package main

import "fmt"

func main() {

	str := "aaabcccdeeef"

	type achar struct {
		char string
		cnt  int
	}

	counts := make(map[int]achar)
	lastIdx := 0
	for i, c := range str {
		fmt.Println(i, string(c))

		found := false
		for idx, ch := range counts {
			if ch.char == string(c) {
				counts[idx].cnt += 1
				fmt.Println("- set  :", string(c), "=", counts[idx].cnt, "[", idx, "]")
				found = true
			}
		}
		if found == false {
			fmt.Println("- added:", string(c), "=", 1, "[", lastIdx, "]")
			counts[lastIdx] = achar{char: string(c), cnt: 1}
			lastIdx++
		}

		//		counts[string(c)] += 1
	}

	fmt.Println("Jetzt untersuchen")
	for i := 0; i <= lastIdx; i++ {
		fmt.Println(i, counts[i].char, counts[i].cnt)
	}

}
