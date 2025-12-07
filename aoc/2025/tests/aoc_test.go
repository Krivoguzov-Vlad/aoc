package aoc2025

import (
	"fmt"
	"testing"

	aoc2025 "github.com/Krivoguzov-Vlad/aoc/aoc/2025"
	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
	"github.com/stretchr/testify/require"
)

func TestAOC(t *testing.T) {
	type answer struct {
		part1 string
		part2 string
	}

	answers := [...]*answer{
		1: {part1: "3", part2: "6"},
		2: {part1: "1227775554", part2: "4174379265"},
		3: {part1: "357", part2: "3121910778619"},
		4: {part1: "13", part2: "43"},
		5: {part1: "3", part2: "14"},
		6: {part1: "4277556", part2: "3263827"},
		7: {part1: "21", part2: "40"},
	}

	for day, answer := range answers {
		if answer == nil {
			continue
		}
		solver := aoc2025.AOC[day]
		if solver == nil {
			continue
		}
		t.Run(fmt.Sprintf("Day %d", day), func(t *testing.T) {
			input := input.MustReadFile(fmt.Sprintf("%d.txt", day))
			solver.ReadInput(input)
			require.Equal(t, answers[day].part1, solver.Part1())
			require.Equal(t, answers[day].part2, solver.Part2())
		})
	}
}
