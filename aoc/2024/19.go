package aoc2024

import (
	"io"
	"strconv"
	"strings"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day19 struct {
	Towels  []string
	Designs []string
}

func (d *Day19) ReadInput(r io.Reader) {
	d.Towels = input.MustReadList[string](r, ", ", "\n")
	input.SkipLine(r)
	d.Designs = input.MustReadList[string](r, "\n")
}

func (d Day19) Part1() string {
	count := 0
	cache := make(map[string]int)
	for _, design := range d.Designs {
		if d.designWays(design, cache) != 0 {
			count++
		}
	}
	return strconv.Itoa(count)
}

func (d Day19) Part2() string {
	count := 0
	cache := make(map[string]int)
	for _, design := range d.Designs {
		count += d.designWays(design, cache)
	}
	return strconv.Itoa(count)
}

func (d Day19) designWays(design string, cache map[string]int) (res int) {
	if design == "" {
		return 1
	}
	defer func() {
		cache[design] = res
	}()
	if v, ok := cache[design]; ok {
		return v
	}
	for _, prefix := range d.Towels {
		if !strings.HasPrefix(design, prefix) {
			continue
		}
		res += d.designWays(design[len(prefix):], cache)
	}
	return res
}
