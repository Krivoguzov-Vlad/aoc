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
	s, e utils.Coordinate

	coordinateToTime map[utils.Coordinate]int
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
	d.coordinateToTime = make(map[utils.Coordinate]int)
	d.m.DFS(d.s, func(depth int, c utils.Coordinate) (needWalk bool) {
		_, ok := d.coordinateToTime[c]
		return !ok && d.m.Get(c) != '#'
	}, func(depth int, c utils.Coordinate) (stop bool) {
		d.coordinateToTime[c] = t
		t++
		return c == d.e
	})
}

func (d *Day20) cheatIter(maxCheatLen int) iter.Seq2[utils.Coordinate, utils.Coordinate] {
	return func(yield func(utils.Coordinate, utils.Coordinate) bool) {
		for cheatStart := range d.coordinateToTime {
			for e := range d.cheatEnds(cheatStart, maxCheatLen) {
				if !yield(cheatStart, e) {
					return
				}
			}
		}
	}
}

func (d *Day20) cheatEnds(c utils.Coordinate, cheatLen int) iter.Seq[utils.Coordinate] {
	return func(yield func(utils.Coordinate) bool) {
		d.cheatEndBFS(c, cheatLen, yield)
	}
}

func (d *Day20) cheatEndBFS(
	s utils.Coordinate,
	maxDepth int,
	yield func(utils.Coordinate) bool,
) {
	d.m.DFS(s, func(depth int, c utils.Coordinate) (needWalk bool) {
		return depth <= maxDepth
	}, func(depth int, c utils.Coordinate) (stop bool) {
		if c == s {
			return false
		}
		_, ok := d.coordinateToTime[c]
		if ok {
			return !yield(c)
		}
		return false
	})
}
