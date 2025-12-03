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
				v, err := ReadValueFromBytes[T](slices.Clone(curValue))
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
			v, err := ReadValueFromBytes[T](curValue)
			_ = yield(v, err)
		}
	}
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
