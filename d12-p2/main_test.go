package main

import "testing"

func TestGetOffspring(t *testing.T) {
	tt := []struct {
		name   string
		pr     PlantRow
		r1, r2 []rune
		m, z   bool
	}{
		{"nmbnz", PlantRow{nil, '#'}, []rune("###z#"), []rune("###.#"), false, false},
		{"mbnz", PlantRow{nil, '#'}, []rune("###z#"), []rune("#####"), true, false},
		{"nmnz", PlantRow{nil, '#'}, []rune("##z##"), []rune("##.##"), false, true},
		{"mz", PlantRow{nil, '#'}, []rune("##z##"), []rune("#####"), true, true},
		{"nmanz", PlantRow{nil, '#'}, []rune("#z###"), []rune("#####"), true, false},
		{"nmaz", PlantRow{nil, '#'}, []rune("#z###"), []rune("#.###"), false, false},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			match, zero := tc.pr.getOffspring(tc.r1, tc.r2)
			if match != tc.m {
				t.Errorf("expected match to be %v got %v", tc.m, match)
			}
			if zero != tc.z {
				t.Errorf("expected zero to be %v got %v", tc.z, zero)
			}
		})
	}
}
