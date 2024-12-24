package input

import (
	"bytes"
	"errors"
	"io"
	"iter"
	"slices"
	"strconv"
)

func MustReadList[T any](r io.Reader, sep string, end ...string) []T {
	res, err := ReadList[T](r, sep, end...)
	if err != nil {
		panic(err)
	}
	return res
}

func ReadList[T any](r io.Reader, sep string, end ...string) ([]T, error) {
	var res []T
	for v, err := range ValueIter[T](r, sep, end...) {
		if err != nil {
			return nil, err
		}
		res = append(res, v)
	}
	return res, nil
}

func ValueIter[T any](r io.Reader, sep string, end ...string) iter.Seq2[T, error] {
	sepBytes := []byte(sep)
	return func(yield func(T, error) bool) {
		var curValue []byte
		for len(end) == 0 || !bytes.HasSuffix(curValue, []byte(end[0])) {
			var b [1]byte
			_, err := r.Read(b[:])
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				var zero T
				_ = yield(zero, err)
				return
			}

			curValue = append(curValue, b[0])

			if bytes.HasSuffix(curValue, sepBytes) {
				curValue = bytes.TrimSuffix(curValue, sepBytes)
				v, err := readValue[T](slices.Clone(curValue))
				if !yield(v, err) {
					return
				}
				curValue = curValue[:0]
			}
		}

		curValue = bytes.TrimSuffix(curValue, sepBytes)
		if len(end) > 0 {
			curValue = bytes.TrimSuffix(curValue, []byte(end[0]))
		}
		if len(curValue) > 0 {
			v, err := readValue[T](curValue)
			_ = yield(v, err)
		}
	}
}

func readValue[T any](data []byte) (T, error) {
	var zero T
	switch v := any(zero).(type) {
	case string:
		return any(string(data)).(T), nil
	case int:
		i, err := strconv.Atoi(string(data))
		if err != nil {
			return zero, err
		}
		return any(i).(T), nil
	case byte:
		if len(data) > 1 {
			panic("too many bytes")
		}
		return any(data[0]).(T), nil
	case Input:
		_, err := v.ReadFrom(bytes.NewReader(data))
		if err != nil {
			return zero, err
		}
		return any(v).(T), nil
	default:
		return zero, errors.New("unknown error")
	}
}

func SkipLine(r io.Reader) {
	for {
		var b [1]byte
		_, err := r.Read(b[:])
		if err != nil {
			panic(err)
		}
		if b[0] == '\n' {
			return
		}
	}
}
