package aoc

import "io"

type Day interface {
	ReadInput(io.Reader)
	Part1() string
	Part2() string
}

type AOC [25]Day
