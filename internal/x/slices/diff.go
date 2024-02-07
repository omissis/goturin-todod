package slices

import (
	"errors"
	"fmt"
	"strconv"

	"golang.org/x/exp/slices"
)

var ErrCannotConvertToString = errors.New("cannot convert value to string")

// Diff returns the added and deleted elements between two slices.
// Type T should implement the fmt.Stringer interface or be a ~string.
// Unfortunately due to https://github.com/golang/go/issues/49054 the two can't be combined in a union type.
func Diff[T any](x, y []T) ([]T, []T, error) {
	d, err := diff(x, y)
	if err != nil {
		return nil, nil, err
	}

	added := make([]T, 0)
	deleted := make([]T, 0)

	for _, v := range d {
		if v[0] != -1 && v[1] == -1 {
			added = append(added, y[v[0]])
		}

		if v[1] != -1 && v[0] == -1 {
			deleted = append(deleted, x[v[1]])
		}
	}

	return slices.Clip(added), slices.Clip(deleted), nil
}

func diff[T any](x, y []T) (map[string][2]int, error) {
	// Diff's values represent which entries were added(index: 0) and removed(index: 1)
	diff := make(map[string][2]int, len(x))

	for k, v := range x {
		vs, err := toString(v)
		if err != nil {
			return nil, err
		}

		diff[vs] = [2]int{-1, k}
	}

	for k, v := range y {
		vs, err := toString(v)
		if err != nil {
			return nil, err
		}

		if entry, ok := diff[vs]; ok {
			entry[0] = k
			diff[vs] = entry

			continue
		}

		diff[vs] = [2]int{k, -1}
	}

	return diff, nil
}

func toString(v any) (string, error) {
	if vs, ok := v.(string); ok {
		return vs, nil
	}

	if vs, ok := v.(fmt.Stringer); ok {
		return vs.String(), nil
	}

	if vi, ok := v.(int); ok {
		return strconv.Itoa(vi), nil
	}

	return "", fmt.Errorf("%w: %v", ErrCannotConvertToString, v)
}
