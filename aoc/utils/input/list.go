package input

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"iter"
	"slices"
	"strconv"
)

func MustReadList[T any](r io.Reader, sep string, opt ...SplitOpt[T]) []T {
	res, err := ReadList[T](r, sep, opt...)
	if err != nil {
		panic(err)
	}
	return res
}

func ReadList[T any](r io.Reader, sep string, opt ...SplitOpt[T]) ([]T, error) {
	var res []T
	for v, err := range ValueIter[T](r, sep, opt...) {
		if err != nil {
			return nil, err
		}
		res = append(res, v)
	}
	return res, nil
}

type SplitOpt[T any] struct {
	Until     string
	WithEmpty bool
	ParseFunc func(io.Reader) (T, error)
}

func ValueIter[T any](r io.Reader, sep string, opt ...SplitOpt[T]) iter.Seq2[T, error] {
	if len(opt) == 0 {
		var zero SplitOpt[T]
		opt = append(opt, zero)
	}

	sepsBytes := [][]byte{[]byte(sep)}
	if len(opt[0].Until) > 0 {
		sepsBytes = append(sepsBytes, []byte(opt[0].Until))
	}

	parseFunc := ReadValueFromBytes[T]
	if opt[0].ParseFunc != nil {
		parseFunc = func(data []byte) (T, error) {
			return opt[0].ParseFunc(bytes.NewReader(data))
		}
	}

	return func(yield func(T, error) bool) {
		var curValue []byte
		var lastSep []byte

		for b, err := range byteIter(r) {
			if err != nil {
				var zero T
				_ = yield(zero, err)
				return
			}

			lastSep = append(lastSep, b)
			curValue = append(curValue, b)
			if len(opt[0].Until) > 0 && len(curValue) <= len(lastSep) {
				if bytes.Equal(lastSep, []byte(opt[0].Until)) {
					return
				}
			}

			if suffix, found := hasAnySuffix(curValue, sepsBytes); found {
				curValue = bytes.TrimSuffix(curValue, suffix)

				if !opt[0].WithEmpty && len(curValue) == 0 {
					continue
				}

				v, err := parseFunc(slices.Clone(curValue))
				if !yield(v, err) {
					return
				}
				curValue = curValue[:0]
			}

			if len(opt[0].Until) > 0 {
				if bytes.Equal(lastSep, []byte(opt[0].Until)) {
					return
				}
				if !bytes.HasPrefix([]byte(opt[0].Until), lastSep) {
					lastSep = lastSep[:0]
				}
				continue
			}
		}

		for _, sep := range sepsBytes {
			curValue = bytes.TrimSuffix(curValue, sep)
		}
		if !opt[0].WithEmpty && len(curValue) == 0 {
			return
		}

		v, err := parseFunc(curValue)
		_ = yield(v, err)
	}
}

func byteIter(r io.Reader) iter.Seq2[byte, error] {
	return func(yield func(byte, error) bool) {
		for {
			var b [1]byte
			_, err := r.Read(b[:])
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				}
				var zero byte
				_ = yield(zero, err)
				return
			}
			if !yield(b[0], nil) {
				return
			}
		}
	}
}

func hasAnySuffix(b []byte, suffixes [][]byte) (suffix []byte, found bool) {
	for _, suffix := range suffixes {
		if bytes.HasSuffix(b, suffix) {
			return suffix, true
		}
	}
	return nil, false
}

func ReadValue[T any](r io.Reader) (T, error) {
	var zero T

	if v, ok := any(&zero).(Input); ok {
		if err := v.ParseFrom(r); err != nil {
			return zero, err
		}
		return *any(v).(*T), nil
	}

	if _, ok := any(zero).(byte); ok {
		var v [1]byte
		_, err := r.Read(v[:])
		if err != nil {
			return zero, err
		}
		return any(v[0]).(T), nil
	}

	data, err := io.ReadAll(r)
	if err != nil {
		return zero, err
	}

	switch any(zero).(type) {
	case string:
		return any(string(data)).(T), nil
	case []byte:
		return any(data).(T), nil
	case int:
		i, err := strconv.Atoi(string(data))
		if err != nil {
			return zero, err
		}
		return any(i).(T), nil
	}

	return zero, fmt.Errorf("type %T %w", zero, errors.ErrUnsupported)
}

func ReadValueFromBytes[T any](data []byte) (T, error) {
	r := bytes.NewReader(data)
	return ReadValue[T](r)
}

func MustReadValue[T any](r io.Reader) T {
	v, err := ReadValue[T](r)
	if err != nil {
		panic(err)
	}
	return v
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
