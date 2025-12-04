package aoc2025

import (
	"io"
	"strconv"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils"
	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day4 struct {
	m *utils.Matrix[byte]
}

func (d *Day4) ReadInput(r io.Reader) {
	d.m = input.MustReadMatrix[byte](r)
}

func (d *Day4) Part1() string {
	count := 0
	for c := range d.m.Iter() {
		if d.canBeGrabbed(c) {
			count++
		}
	}
	return strconv.Itoa(count)
}

func (d *Day4) Part2() string {
	queue := []utils.Cell[byte]{}
	for c := range d.m.Iter() {
		if c.Value() == '@' {
			queue = append(queue, c)
		}
	}

	count := 0
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		if !d.canBeGrabbed(c) {
			continue
		}
		count++
		c.Set('.')
		for n := range c.Neighbours8() {
			if n.Value() == '@' {
				queue = append(queue, n)
			}
		}
	}
	return strconv.Itoa(count)
}

func (d *Day4) canBeGrabbed(c utils.Cell[byte]) bool {
	if c.Value() != '@' {
		return false
	}
	count := 0
	for n := range c.Neighbours8() {
		if n.Value() != '@' {
			continue
		}
		count++
		if count > 3 {
			return false
		}
	}
	return true
}
