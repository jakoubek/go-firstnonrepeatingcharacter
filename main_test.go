package main

import (
	"testing"
)

func TestFindFirstNonRepeatingChar(t *testing.T) {
	var cases = []struct {
		str  string
		want string
	}{
		{"ADBCGHIEFKJLADTVDERFSWVGHQWCNOPENSMSJWIERTFB", "K"},
		{"aaabbcdbd", "c"},
		{"aabbccdd", ""},
	}
	for _, c := range cases {
		t.Run(c.str, func(t *testing.T) {
			have := firstNonRepeatingChar(c.str)
			if have != c.want {
				t.Errorf("got %s, want %s", have, c.want)
			}
		})
	}
}
