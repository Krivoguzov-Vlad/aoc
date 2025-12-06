package aoc2025

import (
	"cmp"
	"io"
	"slices"
	"strconv"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day5 struct {
	ranges []idRange
	ids    []int
}

func (d *Day5) ReadInput(r io.Reader) {
	d.ranges = input.MustReadList[idRange](r, "\n", input.SplitOpt{Until: "\n\n"})
	d.ids = input.MustReadList[int](r, "\n")
	d.mergeRanges()
}

func (d *Day5) Part1() string {
	res := 0
	for _, id := range d.ids {
		_, fresh := slices.BinarySearchFunc(d.ranges, id, func(r idRange, id int) int {
			if r.start <= id && id <= r.end {
				return 0
			}
			return cmp.Compare(r.start, id)
		})
		if fresh {
			res++
		}
	}
	return strconv.Itoa(res)
}

func (d *Day5) Part2() string {
	res := 0
	for _, r := range d.ranges {
		res += r.end - r.start + 1
	}
	return strconv.Itoa(res)
}

func (d *Day5) mergeRanges() {
	slices.SortFunc(d.ranges, func(a, b idRange) int {
		return cmp.Or(
			cmp.Compare(a.start, b.start),
			cmp.Compare(a.end, b.end),
		)
	})

	// would be nice to do it inplace
	mergedRanges := make([]idRange, 0, len(d.ranges))
	mergedRanges = append(mergedRanges, d.ranges[0])
	for _, r := range d.ranges[1:] {
		if mergedRanges[len(mergedRanges)-1].end >= r.start {
			mergedRanges[len(mergedRanges)-1].end = max(mergedRanges[len(mergedRanges)-1].end, r.end)
			continue
		}
		mergedRanges = append(mergedRanges, r)
	}
	d.ranges = mergedRanges
}
