package stacks

import (
	"errors"
	"testing"
)

func TestStackPush(t *testing.T) {
	tests := []struct {
		name      string
		pushItems []int
		wantSizes []int
	}{
		{"push test", []int{1, 4, 9, 7}, []int{1, 2, 3, 4}},
	}
	stack := NewStack[int]()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, pItem := range tt.pushItems {
				stack.push(pItem)
				if stack.len() != tt.wantSizes[i] {
					t.Errorf("Invalid stack size: %v, want %v", stack.len(), tt.wantSizes[i])
				}
			}
		})
	}
}

func TestStackPeekAndPop(t *testing.T) {
	tests := []struct {
		name      string
		pushItems []int
		peekItems []int
	}{
		{"peek test", []int{1, 4, 9, 7}, []int{1, 4, 9, 7}},
	}

	stack := NewStack[int]()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i, pItem := range tt.pushItems {
				stack.push(pItem)
				peek, err := stack.peek()
				errors.Is(err, ErrEmptyStack)
				if peek != tt.peekItems[i] {
					t.Errorf("Invalid stack peek: %v, want %v", peek, tt.peekItems[i])
				}
			}
		})
	}

}

// TestStackPop is a reference for the table-driven shape with error cases.
// Four things to copy: a wantErr column, a fresh stack per row, errors.Is for
// both failure AND success, and asserting the value on the error path too.
func TestStackPop(t *testing.T) {
	tests := []struct {
		name         string
		pushItems    []int // state to build up before the call
		popTimes     int   // how many pops before the one we assert on
		wantVal      int
		wantErr      error // nil means "expect success"
		wantLenAfter int
	}{
		{
			name:    "pop from empty stack",
			wantErr: ErrEmptyStack,
			// wantVal stays 0: pop must return the zero value on failure
		},
		{
			name:         "pop single item",
			pushItems:    []int{42},
			wantVal:      42,
			wantLenAfter: 0,
		},
		{
			name:         "pop returns most recent item",
			pushItems:    []int{1, 4, 9, 7},
			wantVal:      7,
			wantLenAfter: 3,
		},
		{
			name:         "pop after draining errors",
			pushItems:    []int{1, 2},
			popTimes:     2,
			wantErr:      ErrEmptyStack,
			wantLenAfter: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			stack := NewStack[int]() // fresh per row: no state leaks between cases
			for _, item := range tt.pushItems {
				stack.push(item)
			}
			for i := 0; i < tt.popTimes; i++ {
				if _, err := stack.pop(); err != nil {
					t.Fatalf("setup pop #%d: unexpected error %v", i+1, err)
				}
			}

			got, err := stack.pop()

			// errors.Is(nil, nil) is true, so this one check covers both
			// the expect-error and expect-success rows.
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("pop() error = %v, want %v", err, tt.wantErr)
			}
			if got != tt.wantVal {
				t.Errorf("pop() value = %v, want %v", got, tt.wantVal)
			}
			if stack.len() != tt.wantLenAfter {
				t.Errorf("len() after pop = %v, want %v", stack.len(), tt.wantLenAfter)
			}
		})
	}
}
