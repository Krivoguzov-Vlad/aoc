package aoc2025

import (
	"io"
	"strconv"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day3 struct {
	banks [][]byte
}

func (d *Day3) ReadInput(r io.Reader) {
	d.banks = input.MustReadList[[]byte](r, "\n")
	for _, bank := range d.banks {
		for i := range bank {
			bank[i] -= '0'
		}
	}
}

func (d *Day3) Part1() string {
	return d.task(2)
}

func (d *Day3) Part2() string {
	return d.task(12)
}

func (d *Day3) task(n int) string {
	total := 0
	for _, bank := range d.banks {
		joltage := make(joltage, 0, n)
		for i, battery := range bank {
			joltage = joltage.push(battery, len(bank[i:]))
		}
		total += joltage.toInt()
	}
	return strconv.Itoa(total)
}

type joltage []byte

func (jol joltage) push(b byte, batteriesLeft int) joltage {
	if len(jol)+batteriesLeft == cap(jol) {
		return append(jol, b)
	}

	for i := len(jol) - 1; i >= 0; i-- {
		if jol[i] >= b {
			break
		}
		jol = jol[:i]
		if len(jol)+batteriesLeft == cap(jol) {
			break
		}
	}

	if cap(jol) == len(jol) {
		return jol
	}
	return append(jol, b)
}

func (jol joltage) toInt() int {
	res := 0
	for i := range jol {
		res = res*10 + int(jol[i])
	}
	return res
}
