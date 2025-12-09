package aoc2025

import (
	"fmt"
	"io"
	"strconv"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils"
	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day9 struct {
	redTiles []tile
	matrix   *utils.Matrix[byte]
}

func (d *Day9) ReadInput(r io.Reader) {
	d.redTiles = input.MustReadList[tile](r, "\n")
}

func (d *Day9) Part1() string {
	area := 0
	for i, t1 := range d.redTiles {
		for _, t2 := range d.redTiles[i+1:] {
			area = max(area, d.area(t1, t2))
		}
	}
	return strconv.Itoa(area)
}

func (d *Day9) Part2() string {
	area := 0
	for i, t1 := range d.redTiles {
		for _, t2 := range d.redTiles[i+1:] {
			if !d.allGreenOrRed(t1, t2) {
				continue
			}
			area = max(area, d.area(t1, t2))
		}
	}
	return strconv.Itoa(area)
}

func (d *Day9) area(c1, c2 tile) int {
	dx := utils.Abs(c1.X - c2.X)
	dy := utils.Abs(c1.Y - c2.Y)
	return (dx + 1) * (dy + 1)
}

func (d *Day9) allGreenOrRed(c1, c2 tile) bool {
	minX, maxX := min(c1.X, c2.X), max(c1.X, c2.X)
	minY, maxY := min(c1.Y, c2.Y), max(c1.Y, c2.Y)
	for i, current := range append(d.redTiles[1:], d.redTiles[0]) {
		prev := d.redTiles[i]
		mid := utils.Coordinate{X: (prev.X + current.X) / 2, Y: (prev.Y + current.Y) / 2}
		if minX < mid.X && mid.X < maxX && minY < mid.Y && mid.Y < maxY {
			return false
		}
	}
	return true
}

type tile struct {
	utils.Coordinate
}

func (t *tile) ParseFrom(r io.Reader) (err error) {
	_, err = fmt.Fscanf(r, "%d,%d", &t.X, &t.Y)
	return err
}
