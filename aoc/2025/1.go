package aoc2025

import (
	"io"
	"strconv"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils"
	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day1 struct {
	rotates     []rotation
	startArrow  rotation
	numberCount int
}

func (d *Day1) ReadInput(r io.Reader) {
	d.numberCount = 100
	d.startArrow = rotation(50)
	for v, err := range input.ValueIter[rotation](r, "\n") {
		if err != nil {
			panic(err)
		}
		d.rotates = append(d.rotates, v)
	}
}

func (d *Day1) Part1() string {
	arrow := d.startArrow
	res := 0
	for _, rot := range d.rotates {
		arrow = (arrow + rot) % rotation(d.numberCount)
		if arrow == 0 {
			res++
		}
	}
	return strconv.Itoa(res)
}

func (d *Day1) Part2() string {
	arrow := d.startArrow
	res := 0
	for _, rot := range d.rotates {
		arrow += rot

		res += utils.Abs(int(arrow) / d.numberCount)
		if arrow <= 0 && arrow != rot {
			res++
		}

		arrow = arrow % rotation(d.numberCount)
		if arrow < 0 {
			arrow += rotation(d.numberCount)
		}
	}
	return strconv.Itoa(res)
}

type rotation int

func (rot *rotation) ParseFrom(r io.Reader) (err error) {
	direction := input.MustReadValue[byte](r)
	*rot = rotation(input.MustReadValue[int](r))
	if direction == 'L' {
		*rot = -*rot
	}
	return nil
}
