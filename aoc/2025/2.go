package aoc2025

import (
	"fmt"
	"io"
	"strconv"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day2 struct {
	ranges []idRange
}

func (d *Day2) ReadInput(r io.Reader) {
	d.ranges = input.MustReadList[idRange](r, ",", "\n")
}

func (d *Day2) Part1() string {
	var res int64
	for _, r := range d.ranges {
		for i := r.start; i <= r.end; i++ {
			s := strconv.Itoa(i)
			if d.containsRepeated(s, 2) {
				res += int64(i)
			}
		}
	}
	return strconv.FormatInt(res, 10)
}

func (d *Day2) Part2() string {
	var res int64
	for _, r := range d.ranges {
		for i := r.start; i <= r.end; i++ {
			s := strconv.Itoa(i)
			for j := 2; j <= len(s); j++ {
				if d.containsRepeated(s, j) {
					res += int64(i)
					break
				}
			}
		}
	}
	return strconv.FormatInt(res, 10)
}

func (d *Day2) containsRepeated(s string, repeatedCount int) bool {
	if len(s)%repeatedCount != 0 {
		return false
	}
	partLen := len(s) / repeatedCount
	for i := 0; i < len(s); i += partLen {
		if s[:partLen] != s[i:i+partLen] {
			return false
		}
	}
	return true
}

type idRange struct {
	start int
	end   int
}

func (ir *idRange) ParseFrom(r io.Reader) (err error) {
	_, err = fmt.Fscanf(r, "%d-%d", &ir.start, &ir.end)
	return err
}
