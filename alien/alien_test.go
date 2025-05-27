package alien

import "testing"

func TestAlienToInt(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{"Simple AAA", "AAA", 3},
		{"LBAAA", "LBAAA", 58},       
		{"RCRZCAB", "RCRZCAB", 1994}, 
		{"Subtractive AB = 4", "AB", 4},
		{"Subtractive AZ = 9", "AZ", 9},
		{"Subtractive ZL = 40", "ZL", 40},
		{"Subtractive ZC = 90", "ZC", 90},
		{"Subtractive CD = 400", "CD", 400},
		{"Subtractive CR = 900", "CR", 900},
		{"Complex case", "RZLCBAAA", 1164}, 
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := AlienToInt(tt.input)
			if result != tt.expected {
				t.Errorf("expected %d but got %d", tt.expected, result)
			}
		})
	}
}
