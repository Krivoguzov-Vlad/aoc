package aoc2024

import (
	"io"
	"iter"
	"strconv"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils"
	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day20 struct {
	m    *utils.Matrix[byte]
	s, e utils.Cell[byte]

	coordinateToTime map[utils.Cell[byte]]int
}

func (d *Day20) ReadInput(r io.Reader) {
	d.m = input.MustReadMatrix[byte](r)
	d.s = d.m.Find('S')
	d.e = d.m.Find('E')

}

func (d *Day20) Part1() string {
	return strconv.Itoa(d.countCheats(2, 100))
}

func (d *Day20) Part2() string {
	return strconv.Itoa(d.countCheats(20, 100))
}

func (d *Day20) countCheats(maxCheatLen int, minSave int) int {
	d.fillTimes()
	count := 0
	for start, end := range d.cheatIter(maxCheatLen) {
		t1 := d.coordinateToTime[start]
		t2 := d.coordinateToTime[end]
		diff := t2 - t1
		diff -= utils.Abs(start.X-end.X) + utils.Abs(start.Y-end.Y)
		if diff >= minSave {
			count++
		}
	}
	return count
}

func (d *Day20) fillTimes() {
	t := 0
	d.coordinateToTime = make(map[utils.Cell[byte]]int)
	for _, cell := range d.s.DFS(func(depth int, c utils.Cell[byte]) (needWalk bool) {
		_, ok := d.coordinateToTime[c]
		return !ok && c.Value() != '#'
	}) {
		d.coordinateToTime[cell] = t
		t++
		if cell == d.e {
			return
		}
	}
}

func (d *Day20) cheatIter(maxCheatLen int) iter.Seq2[utils.Cell[byte], utils.Cell[byte]] {
	return func(yield func(utils.Cell[byte], utils.Cell[byte]) bool) {
		for cheatStart := range d.coordinateToTime {
			for e := range d.cheatEnds(cheatStart, maxCheatLen) {
				if !yield(cheatStart, e) {
					return
				}
			}
		}
	}
}

func (d *Day20) cheatEnds(c utils.Cell[byte], cheatLen int) iter.Seq[utils.Cell[byte]] {
	return func(yield func(utils.Cell[byte]) bool) {
		d.cheatEndBFS(c, cheatLen, yield)
	}
}

func (d *Day20) cheatEndBFS(
	s utils.Cell[byte],
	maxDepth int,
	yield func(utils.Cell[byte]) bool,
) {
	for _, cell := range s.DFS(func(depth int, c utils.Cell[byte]) (needWalk bool) {
		return depth <= maxDepth
	}) {
		_, ok := d.coordinateToTime[cell]
		if ok && !yield(cell) {
			return
		}
	}
}
