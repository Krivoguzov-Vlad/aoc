package aoc2025

import (
	"cmp"
	"fmt"
	"io"
	"maps"
	"slices"
	"sort"
	"strconv"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day8 struct {
	points      []point
	connections []connection
}

func (d *Day8) ReadInput(r io.Reader) {
	d.points = input.MustReadList[point](r, "\n")
	d.connections = make([]connection, 0, len(d.points)*(len(d.points)))
	for i := range d.points {
		p1 := d.points[i]
		for j := range d.points[i+1:] {
			j += i + 1
			p2 := d.points[j]
			d.connections = append(d.connections, connection{
				point1Idx: i,
				point2Idx: j,
				distance:  p1.Distance(p2),
			})
		}
	}
	slices.SortFunc(d.connections, func(a, b connection) int {
		return cmp.Compare(a.distance, b.distance)
	})
}

func (d *Day8) Part1() string {
	connectionLimit := 1000
	if len(d.points) == 20 {
		// example input
		connectionLimit = 10
	}

	setUnion := newSetUnion(len(d.points))
	for _, conn := range d.connections[:connectionLimit] {
		setUnion.mergeElemSets(conn.point1Idx, conn.point2Idx)
	}

	elemsCount := make(map[int]int)
	for i := range setUnion.elemToSet {
		elemsCount[setUnion.getElemSet(i)]++
	}
	counts := slices.Collect(maps.Values(elemsCount))
	sort.Ints(counts)

	res := 1
	for _, count := range counts[len(elemsCount)-3:] {
		res *= count
	}
	return strconv.Itoa(res)
}

func (d *Day8) Part2() string {
	setUnion := newSetUnion(len(d.points))
	connectedCount := 0
	for _, conn := range d.connections {
		if !setUnion.mergeElemSets(conn.point1Idx, conn.point2Idx) {
			continue
		}
		connectedCount++
		if connectedCount == len(d.points)-1 {
			return strconv.Itoa(d.points[conn.point1Idx].x * d.points[conn.point2Idx].x)
		}
	}
	return "unreachable"
}

type point struct {
	x, y, z int
}

func (p *point) ParseFrom(r io.Reader) (err error) {
	_, err = fmt.Fscanf(r, "%d,%d,%d", &p.x, &p.y, &p.z)
	return err
}

func (p point) Distance(o point) int {
	dx := p.x - o.x
	dy := p.y - o.y
	dz := p.z - o.z

	return dx*dx + dy*dy + dz*dz
}

type connection struct {
	point1Idx, point2Idx int
	distance             int
}

type setUnion struct {
	elemToSet []int
}

func newSetUnion(n int) *setUnion {
	res := &setUnion{elemToSet: make([]int, n)}
	// initially each element is in its own set
	for i := range res.elemToSet {
		res.elemToSet[i] = i
	}
	return res
}

func (s *setUnion) getElemSet(elem int) int {
	set := s.elemToSet[elem]
	if set == elem {
		return set
	}
	set = s.getElemSet(set)
	s.elemToSet[elem] = set
	return set
}

func (s *setUnion) mergeElemSets(elem1, elem2 int) bool {
	set1 := s.getElemSet(elem1)
	set2 := s.getElemSet(elem2)
	if set1 == set2 {
		return false
	}

	// can we do it better than O(n)?
	for i, v := range s.elemToSet {
		if v == set2 {
			s.elemToSet[i] = set1
		}
	}
	return true
}
