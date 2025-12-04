package utils

import "iter"

type Cell[T comparable] struct {
	c Coordinate
	m *Matrix[T]
}

func (mv *Cell[T]) Set(v T) {
	mv.m.Set(mv.c, v)
}

func (mv Cell[T]) Value() T {
	return mv.m.Get(mv.c)
}

// only valid neighbours
func (mv Cell[T]) Neighbours() iter.Seq[Cell[T]] {
	return func(yield func(Cell[T]) bool) {
		for _, n := range mv.c.Neighbours() {
			if !mv.m.IsValid(n) {
				continue
			}
			if !yield(Cell[T]{c: n, m: mv.m}) {
				return
			}
		}
	}
}

func (mv Cell[T]) Neighbours8() iter.Seq[Cell[T]] {
	return func(yield func(Cell[T]) bool) {
		for _, n := range mv.c.Neighbours8() {
			if !mv.m.IsValid(n) {
				continue
			}
			if !yield(Cell[T]{c: n, m: mv.m}) {
				return
			}
		}
	}
}
