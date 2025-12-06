package utils

func Abs[T int](a T) T {
	if a < 0 {
		return -a
	}
	return a
}

func Transpose[T comparable](m [][]T) [][]T {
	res := make([][]T, len(m[0]))
	for i := range res {
		res[i] = make([]T, len(m))
	}
	for i := range m {
		for j := range m[i] {
			res[j][i] = m[i][j]
		}
	}
	return res
}
