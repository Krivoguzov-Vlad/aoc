package input

import (
	"bufio"
	"bytes"
	"io"

	"github.com/Krivoguzov-Vlad/aoc/aoc/utils"
)

func MustReadMatrix[T comparable](r io.Reader, sep ...string) *utils.Matrix[T] {
	m, err := ReadMatrix[T](r, sep...)
	if err != nil {
		panic(err)
	}
	return m
}

func ReadMatrix[T comparable](r io.Reader, sep ...string) (*utils.Matrix[T], error) {
	if len(sep) == 0 {
		sep = append(sep, "")
	}

	scanner := bufio.NewScanner(r)
	var m utils.Matrix[T]
	for scanner.Scan() {
		var row []T
		values := ValueIter[T](bytes.NewReader(scanner.Bytes()), sep[0])
		for v, err := range values {
			if err != nil {
				return nil, err
			}
			row = append(row, v)
		}
		if len(row) > 0 {
			m.AddRow(row)
		}
	}
	return &m, nil
}
