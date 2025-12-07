package aoc2025

import (
	"io"
	"strconv"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils"
	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day7 struct {
	m *utils.Matrix[byte]
	s utils.Cell[byte]
}

func (d *Day7) ReadInput(r io.Reader) {
	d.m = input.MustReadMatrix[byte](r)
	d.s = d.m.Find('S')
	d.s.Set('.')
}

func (d *Day7) Part1() string {
	return strconv.Itoa(d.countSplits(d.s, make(map[utils.Cell[byte]]bool)))
}

func (d *Day7) Part2() string {
	return strconv.Itoa(1 + d.countTimelines(d.s, make(map[utils.Cell[byte]]int)))
}

func (d *Day7) countSplits(cell utils.Cell[byte], visited map[utils.Cell[byte]]bool) int {
	if !cell.IsValid() {
		return 0
	}
	if _, ok := visited[cell]; ok {
		return 0
	}
	visited[cell] = true
	switch cell.Value() {
	case '.':
		return d.countSplits(cell.Down(), visited)
	case '^':
		return 1 + d.countSplits(cell.Left(), visited) + d.countSplits(cell.Right(), visited)
	default:
		return 0
	}
}

func (d *Day7) countTimelines(cell utils.Cell[byte], cache map[utils.Cell[byte]]int) (res int) {
	if !cell.IsValid() {
		return 0
	}
	if v, ok := cache[cell]; ok {
		return v
	}
	defer func() {
		cache[cell] = res
	}()
	switch cell.Value() {
	case '.':
		return d.countTimelines(cell.Down(), cache)
	case '^':
		return 1 + d.countTimelines(cell.Left(), cache) + d.countTimelines(cell.Right(), cache)
	default:
		return 0
	}
}
