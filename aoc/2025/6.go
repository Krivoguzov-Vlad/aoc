package aoc2025

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils/input"
)

type Day6 struct {
	ops   []operation
	part1 [][]int
	part2 [][]int
}

func (d *Day6) ReadInput(r io.Reader) {
	data, _ := io.ReadAll(r)
	data = bytes.Trim(data, "\n")

	idx := bytes.LastIndex(data, []byte("\n"))
	numbersData := data[:idx]
	opsData := data[idx+1:]

	d.ops = input.MustReadList[operation](bytes.NewReader(opsData), " ")
	d.part1 = input.MustReadMatrix[int](bytes.NewReader(numbersData), " ").
		Transpose().Values

	tmp := input.MustReadMatrix[byte](bytes.NewReader(numbersData)).Transpose()
	for i, row := range tmp.Values {
		tmp.Values[i] = bytes.TrimSpace(row)
	}
	numbersPart2Input := bytes.NewReader(bytes.Join(tmp.Values, []byte("\n")))
	d.part2 = input.MustReadList(numbersPart2Input, "\n\n", input.SplitOpt[[]int]{
		ParseFunc: func(r io.Reader) ([]int, error) {
			return input.ReadList[int](r, "\n")
		},
	})
}

func (d *Day6) Part1() string {
	return strconv.Itoa(d.solve(d.ops, d.part1))
}

func (d *Day6) Part2() string {
	return strconv.Itoa(d.solve(d.ops, d.part2))
}

func (d *Day6) solve(ops []operation, problems [][]int) int {
	total := 0
	for i, problem := range problems {
		answer := problem[0]
		op := ops[i]
		for _, number := range problem[1:] {
			answer = op.apply(answer, number)
		}
		total += answer
	}
	return total
}

type operation byte

func (o *operation) ParseFrom(r io.Reader) error {
	v, err := input.ReadValue[byte](r)
	if err != nil {
		return err
	}
	*o = operation(v)
	return nil
}

func (op operation) apply(a, b int) int {
	switch op {
	case '*':
		return a * b
	case '+':
		return a + b
	default:
		panic(fmt.Errorf("operation: %w", errors.ErrUnsupported))
	}
}
