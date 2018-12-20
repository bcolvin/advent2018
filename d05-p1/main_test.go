package main

import "testing"

func TestReact(t *testing.T) {
	tt := []struct {
		in, out string
		ok      bool
	}{
		{"aA", "", false},
		{"AaA", "A", false},
		{"AaABbB", "AB", false},
		{"AaACcBbB", "AB", false},
		{"aa", "aa", true},
		{"Baab", "Baab", true},
		{"dabAcCaCBAcCcaDA", "dabAaCBAcaDA", false},
		{"dabAaCBAcaDA", "dabCBAcaDA", false},
		{"dabCBAcaDA", "dabCBAcaDA", true},
	}
	for _, tc := range tt {
		t.Run(tc.in, func(t *testing.T) {
			out, ok := react([]rune(tc.in))
			if tc.ok != ok {
				t.Errorf("expected ok to be %v got %v", tc.ok, ok)
			}
			outStr := string(out)
			if tc.out != outStr {
				t.Errorf("expected out to be %v got %v", tc.out, outStr)
			}
		})
	}
}
