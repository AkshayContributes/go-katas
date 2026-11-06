package stacks

import (
	"errors"
)

var ErrEmptyStack = errors.New("stack is empty")

type node[T any] struct {
	val  T
	next *node[T]
	prev *node[T]
}

// type Stack string
type Stack[T any] struct {
	root *node[T]
	top  *node[T]
	size int
}

func NewStack[T any]() *Stack[T] {

	root := &node[T]{}

	stack := Stack[T]{
		root: root,
		top:  root,
		size: 0,
	}

	return &stack

}

func (s *Stack[T]) pop() (T, error) {
	var zero T
	if s.size == 0 {
		return zero, ErrEmptyStack
	}

	top := s.top.val
	s.top = s.top.prev
	s.size -= 1
	return top, nil
}

func (s *Stack[T]) peek() (T, error) {
	var zero T
	if s.size == 0 {
		return zero, ErrEmptyStack
	}

	top := s.top.val
	return top, nil
}

func (s *Stack[T]) push(val T) {
	newNode := &node[T]{
		val:  val,
		prev: s.top,
		next: nil,
	}
	s.top.next = newNode
	s.top = s.top.next
	s.size += 1
}

func (s *Stack[T]) len() int {
	return s.size
}
