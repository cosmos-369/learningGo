package generics

type Stack[T any] []T

func (s Stack[T]) IsEmpty() bool {
	return len(s) == 0
}

func (s *Stack[T]) Push(val T) {
	*s = append(*s, val)
}

func (s *Stack[T]) Pop() (T, bool) {

	if s.IsEmpty() {
		var zero T
		return zero, false
	}

	len := len(*s)
	val := (*s)[len-1]
	*s = (*s)[:(len - 1)]
	return val, true
}
