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

func (m Matrix[T]) Find(item T) Cell[T] {
	return m.FindFunc(func(cell Cell[T]) bool {
		return cell.Value() == item
	})
}

func (m *Matrix[T]) FindFunc(f func(cell Cell[T]) bool) Cell[T] {
	for i := range m.Values {
		for j := range m.Values[i] {
			cell := Cell[T]{Coordinate: Coordinate{X: j, Y: i}, m: m}
			if f(cell) {
				return cell
			}
		}
	}
	return Cell[T]{Coordinate: Coordinate{X: -1, Y: -1}, m: m}
}

func (m *Matrix[T]) Iter() iter.Seq[Cell[T]] {
	return func(yield func(Cell[T]) bool) {
		for i := range m.Values {
			for j := range m.Values[i] {
				if !yield(Cell[T]{Coordinate: Coordinate{X: j, Y: i}, m: m}) {
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

func (m *Matrix[T]) DFS(start Coordinate, needWalk func(depth int, c Cell[T]) bool) iter.Seq2[int, Cell[T]] {
	var queue []Coordinate
	queue = append(queue, start)
	addedToQueue := make(map[Coordinate]bool)

	return func(yield func(int, Cell[T]) bool) {
		for depth := 0; len(queue) > 0; depth++ {
			layer := queue
			queue = queue[len(layer):]
			for _, item := range layer {
				cell := Cell[T]{Coordinate: item, m: m}
				if !yield(depth, cell) {
					return
				}

				for c := range cell.Neighbours() {
					if addedToQueue[c.Coordinate] {
						continue
					}
					if !needWalk(depth+1, c) {
						continue
					}
					queue = append(queue, c.Coordinate)
					addedToQueue[c.Coordinate] = true
				}
			}
		}
	}
}
