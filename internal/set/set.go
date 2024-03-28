package set

import "fmt"

type Set[T comparable] map[T]struct{}

func NewSet[T comparable](values ...T) Set[T] {
	set := make(Set[T])

	for _, value := range values {
		set[value] = struct{}{}
	}

	return set
}

func (set Set[T]) Add(values ...T) {
	for _, value := range values {
		set[value] = struct{}{}
	}
}

func (set Set[T]) Delete(values ...T) {
	for _, value := range values {
		delete(set, value)
	}
}

func (set Set[T]) Clear() {
	for value := range set {
		delete(set, value)
	}
}

func (set Set[T]) Size() int {
	return len(set)
}

func (set Set[T]) IsEmpty() bool {
	return len(set) == 0
}

func (set Set[T]) Contains(value T) bool {
	_, ok := set[value]
	return ok
}

func (set Set[T]) ToSlice() []T {
	values := make([]T, 0, len(set))

	for value := range set {
		values = append(values, value)
	}

	return values
}

func (s Set[T]) String() string {
	return fmt.Sprintf("%v", s.ToSlice())
}
