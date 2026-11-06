package reverseslice

import (
	"slices"
	"testing"
)

func TestReverseInts(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"single", []int{1}, []int{1}},
		{"even length", []int{1, 2, 3, 4}, []int{4, 3, 2, 1}},
		{"odd length", []int{1, 2, 3}, []int{3, 2, 1}},
		{"duplicates", []int{7, 7, 8}, []int{8, 7, 7}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Reverse(tt.in)
			if !slices.Equal(got, tt.want) {
				t.Errorf("Reverse(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

func TestReverseStrings(t *testing.T) {
	in := []string{"a", "b", "c"}
	want := []string{"c", "b", "a"}

	if got := Reverse(in); !slices.Equal(got, want) {
		t.Errorf("Reverse(%v) = %v, want %v", in, got, want)
	}
}

func TestReverseDoesNotMutateInput(t *testing.T) {
	in := []int{1, 2, 3}
	original := slices.Clone(in)

	Reverse(in)

	if !slices.Equal(in, original) {
		t.Errorf("input was mutated: got %v, want %v", in, original)
	}
}

func TestReverseNilStaysNil(t *testing.T) {
	if got := Reverse[int](nil); got != nil {
		t.Errorf("Reverse(nil) = %v, want nil", got)
	}
}

func BenchmarkReverse(b *testing.B) {
	in := make([]int, 1024)
	for i := range in {
		in[i] = i
	}

	for b.Loop() {
		Reverse(in)
	}
}
