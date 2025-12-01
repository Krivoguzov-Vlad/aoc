package input

import (
	"bytes"
	"io"
	"os"
)

type Input interface {
	ParseFrom(r io.Reader) (err error)
}

func MustReadFile(path string) io.Reader {
	data, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	return bytes.NewReader(data)
}
