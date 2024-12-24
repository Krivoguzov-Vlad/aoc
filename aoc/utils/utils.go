package utils

func Abs[T int](a T) T {
	if a < 0 {
		return -a
	}
	return a
}
