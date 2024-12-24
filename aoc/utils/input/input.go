package input

import (
	"io"
)

type Input interface {
	io.ReaderFrom
}
