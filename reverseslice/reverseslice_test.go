package reverseslice

import (
	"slices"
	"testing"
)

// One worked subtest as a template. Fill in the rest of the cases:
// nil, empty, single element, even length, odd length, duplicates.
func TestReverseInts(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"odd length", []int{1, 2, 3}, []int{3, 2, 1}},
		{"even length", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotManual := ReverseManual(tt.in)
			if !slices.Equal(gotManual, tt.want) {
				t.Errorf("Reverse(%v) = %v, want %v", tt.in, gotManual, tt.want)
			}

			gotSlice := ReverseUsingSlice(tt.in)
			if !slices.Equal(gotSlice, tt.want) {
				t.Errorf("Reverse(%v) = %v, want %v", tt.in, gotSlice, tt.want)
			}
		})
	}
}

func TestReverseStrings(t *testing.T) {
	tests := []struct {
		name  string
		input []string
		want  []string
	}{
		{"string slice", []string{"Alpha", "Beta", "Cheetah"}, []string{"Cheetah", "Beta", "Alpha"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReverseManual(tt.input)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Reverse(%v) = %v, want %v", tt.input, got, tt.want)
			}

			gotSlice := ReverseUsingSlice(tt.input)
			if !slices.Equal(gotSlice, tt.want) {
				t.Errorf("Reverse(%v) = %v, want %v", tt.input, gotSlice, tt.want)
			}
		})
	}
}

// TODO: TestReverseStrings          - exercise the generic with a second type
// TODO: TestReverseDoesNotMutate    - clone the input, call Reverse, compare
// TODO: BenchmarkReverse            - see the b.Loop() note
