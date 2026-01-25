package gen

import (
	"testing"
)

func TestRandom(t *testing.T) {
	tests := []struct {
		name    string
		wantMin int
		wantMax int
	}{
		{"digits 1", 1, 10},
		{"digits 2", 10, 100},
		{"digits 3", 100, 1000},
		{"digits 4", 1000, 10000},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Random(tt.wantMin, tt.wantMax)
			if got < tt.wantMin || got >= tt.wantMax {
				t.Errorf("Random() = %v, want between %v and %v", got, tt.wantMin, tt.wantMax)
			}
		})
	}
}
