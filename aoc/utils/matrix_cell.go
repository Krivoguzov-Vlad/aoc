package utils

import "iter"

type Cell[T comparable] struct {
	Coordinate
	m *Matrix[T]
}

func (mv *Cell[T]) Set(v T) {
	mv.m.Set(mv.Coordinate, v)
}

func (mv Cell[T]) Value() T {
	return mv.m.Get(mv.Coordinate)
}

// only valid neighbours
func (mv Cell[T]) Neighbours() iter.Seq[Cell[T]] {
	return func(yield func(Cell[T]) bool) {
		for _, n := range mv.Coordinate.Neighbours() {
			if !mv.m.IsValid(n) {
				continue
			}
			if !yield(Cell[T]{Coordinate: n, m: mv.m}) {
				return
			}
		}
	}
}

func (mv Cell[T]) Neighbours8() iter.Seq[Cell[T]] {
	return func(yield func(Cell[T]) bool) {
		for _, n := range mv.Coordinate.Neighbours8() {
			if !mv.m.IsValid(n) {
				continue
			}
			if !yield(Cell[T]{Coordinate: n, m: mv.m}) {
				return
			}
		}
	}
}

func (v Cell[T]) DFS(needWalk func(depth int, c Cell[T]) bool) iter.Seq2[int, Cell[T]] {
	return v.m.DFS(v.Coordinate, needWalk)
}
