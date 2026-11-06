package countwords

import (
	"maps"
	"testing"
)

func TestCountWordsInString(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want int
	}{
		{"formatted string", "quick brown fox jumps over me", 6},
		{"ill formatted string", "     quick   brown    fox ", 3},
		{"only whitespace", "           ", 0},
		{"only newline", "\n", 0},
		{"empty string", "", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWordsInString(tt.in)
			if got != tt.want {
				t.Errorf("CountWordsInString(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestCountWordFrequency(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want map[string]int
	}{
		{"simple string", "      quick quick brown brown brown fox", map[string]int{"quick": 2, "brown": 3, "fox": 1}},
		{"multiline string", "quick\nquick", map[string]int{"quick": 2}},
		{"trailing newline string", "quick quick\n", map[string]int{"quick": 2}},
		{"CRLF string", "a b\r\na b", map[string]int{"a": 2, "b": 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CountWordFrequency(tt.in)
			if !maps.Equal(got, tt.want) {
				t.Errorf("CountWordFrequency(%q): %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
