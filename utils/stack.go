package utils

type Stack[T any] struct {
	elements []T
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}

func (s *Stack[T]) Pop() T {
	var pop T
	if len(s.elements) == 0 {
		return pop
	}
	pop = s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return pop
}

func (s *Stack[T]) Top() T {
	var top T
	if len(s.elements) == 0 {
		return top
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) Size() int {
	return len(s.elements)
}
