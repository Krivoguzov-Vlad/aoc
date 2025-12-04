package utils

import (
	"fmt"
	"iter"
)

type Matrix[T comparable] struct {
	Values [][]T
}

func (m *Matrix[T]) AddRow(row []T) {
	m.Values = append(m.Values, row)
}

func (m Matrix[T]) IsValid(c Coordinate) bool {
	return 0 <= c.X && c.X < len(m.Values[0]) &&
		0 <= c.Y && c.Y < len(m.Values)
}

func (m Matrix[T]) Set(c Coordinate, v T) {
	m.Values[c.Y][c.X] = v
}

func (m Matrix[T]) Get(c Coordinate) T {
	return m.Values[c.Y][c.X]
}

func (m Matrix[T]) Find(item T) Coordinate {
	return m.FindFunc(func(t T) bool {
		return t == item
	})
}

func (m Matrix[T]) FindFunc(f func(T) bool) Coordinate {
	for i := range m.Values {
		for j := range m.Values[i] {
			if f(m.Values[i][j]) {
				return Coordinate{X: j, Y: i}
			}
		}
	}
	return Coordinate{X: -1, Y: -1}
}

func (m Matrix[T]) Iter() iter.Seq[Cell[T]] {
	return func(yield func(Cell[T]) bool) {
		for i := range m.Values {
			for j := range m.Values[i] {
				if !yield(Cell[T]{c: Coordinate{X: j, Y: i}, m: &m}) {
					return
				}
			}
		}
	}
}

func (m Matrix[T]) Print() {
	for i := range m.Values {
		for j := range m.Values[i] {
			v := m.Values[i][j]
			switch any(v).(type) {
			case byte:
				fmt.Printf("%c", v)
			default:
				fmt.Printf("%v", v)
			}
		}
		fmt.Println()
	}
}

func (m Matrix[T]) DFS(
	s Coordinate,
	needWalk func(depth int, c Coordinate) (needWalk bool),
	foreach func(depth int, c Coordinate) (stop bool)) {
	var queue []Coordinate
	queue = append(queue, s)
	addedToQueue := make(map[Coordinate]bool)

	for depth := 0; len(queue) > 0; depth++ {
		layer := queue
		queue = queue[len(layer):]
		for _, item := range layer {
			if foreach(depth, item) {
				return
			}

			for _, c := range item.Neighbours() {
				if !m.IsValid(c) {
					continue
				}
				if addedToQueue[c] {
					continue
				}
				if !needWalk(depth+1, c) {
					continue
				}
				queue = append(queue, c)
				addedToQueue[c] = true
			}
		}
	}
}
