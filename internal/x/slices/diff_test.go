package slices_test

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"

	"github.com/omissis/goturin-todod/internal/x/slices"
)

type ex struct {
	val string
}

func (e ex) String() string {
	return e.val
}

func TestStrDiff(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		x       []string
		y       []string
		wantAdd []string
		wantDel []string
		wantErr error
	}{
		{
			desc:    "slices are empty",
			x:       []string{},
			y:       []string{},
			wantAdd: []string{},
			wantDel: []string{},
		},
		{
			desc:    "slices are equal",
			x:       []string{"a", "b", "c"},
			y:       []string{"a", "b", "c"},
			wantAdd: []string{},
			wantDel: []string{},
		},
		{
			desc:    "value 'c' is deleted",
			x:       []string{"a", "b", "c"},
			y:       []string{"a", "b"},
			wantAdd: []string{},
			wantDel: []string{"c"},
		},
		{
			desc:    "value 'd' is added",
			x:       []string{"a", "b", "c"},
			y:       []string{"a", "b", "c", "d"},
			wantAdd: []string{"d"},
			wantDel: []string{},
		},
		{
			desc:    "value 'c' is deleted, value 'd' is added",
			x:       []string{"a", "b", "c"},
			y:       []string{"a", "b", "d"},
			wantAdd: []string{"d"},
			wantDel: []string{"c"},
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			added, deleted, err := slices.Diff(tC.x, tC.y)

			if !cmp.Equal(added, tC.wantAdd, cmpopts.EquateEmpty()) {
				t.Errorf("wantAdd = %v, got = %v", tC.wantAdd, added)
			}

			if !cmp.Equal(deleted, tC.wantDel, cmpopts.EquateEmpty()) {
				t.Errorf("wantDel = %v, got = %v", tC.wantDel, deleted)
			}

			if !cmp.Equal(err, tC.wantErr, cmpopts.EquateErrors()) {
				t.Errorf("wantErr = %v, got = %v", tC.wantErr, err)
			}
		})
	}
}

func TestStringerDiff(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		x       []ex
		y       []ex
		wantAdd []ex
		wantDel []ex
		wantErr error
	}{
		{
			desc:    "slices are empty",
			x:       []ex{},
			y:       []ex{},
			wantAdd: []ex{},
			wantDel: []ex{},
		},
		{
			desc:    "slices are equal",
			x:       []ex{{"a"}, {"b"}, {"c"}},
			y:       []ex{{"a"}, {"b"}, {"c"}},
			wantAdd: []ex{},
			wantDel: []ex{},
		},
		{
			desc:    "value 'c' is deleted",
			x:       []ex{{"a"}, {"b"}, {"c"}},
			y:       []ex{{"a"}, {"b"}},
			wantAdd: []ex{},
			wantDel: []ex{{"c"}},
		},
		{
			desc:    "value '4' is added",
			x:       []ex{{"a"}, {"b"}, {"c"}},
			y:       []ex{{"a"}, {"b"}, {"c"}, {"d"}},
			wantAdd: []ex{{"d"}},
			wantDel: []ex{},
		},
		{
			desc:    "value 'c' is deleted, value 'd' is added",
			x:       []ex{{"a"}, {"b"}, {"c"}},
			y:       []ex{{"a"}, {"b"}, {"d"}},
			wantAdd: []ex{{"d"}},
			wantDel: []ex{{"c"}},
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			added, deleted, err := slices.Diff(tC.x, tC.y)

			if !cmp.Equal(added, tC.wantAdd, cmpopts.EquateEmpty(), cmp.AllowUnexported(ex{})) {
				t.Errorf("wantAdd = %v, got = %v", tC.wantAdd, added)
			}

			if !cmp.Equal(deleted, tC.wantDel, cmpopts.EquateEmpty(), cmp.AllowUnexported(ex{})) {
				t.Errorf("wantDel = %v, got = %v", tC.wantDel, deleted)
			}

			if !cmp.Equal(err, tC.wantErr, cmpopts.EquateErrors()) {
				t.Errorf("wantErr = %v, got = %v", tC.wantErr, err)
			}
		})
	}
}

func TestIntDiff(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		desc    string
		x       []int
		y       []int
		wantAdd []int
		wantDel []int
		wantErr error
	}{
		{
			desc:    "slices are empty",
			x:       []int{},
			y:       []int{},
			wantAdd: []int{},
			wantDel: []int{},
		},
		{
			desc:    "slices are equal",
			x:       []int{1, 2, 3},
			y:       []int{1, 2, 3},
			wantAdd: []int{},
			wantDel: []int{},
		},
		{
			desc:    "value '3' is deleted",
			x:       []int{1, 2, 3},
			y:       []int{1, 2},
			wantAdd: []int{},
			wantDel: []int{3},
		},
		{
			desc:    "value '4' is added",
			x:       []int{1, 2, 3},
			y:       []int{1, 2, 3, 4},
			wantAdd: []int{4},
			wantDel: []int{},
		},
		{
			desc:    "value '3' is deleted, value '4' is added",
			x:       []int{1, 2, 3},
			y:       []int{1, 2, 4},
			wantAdd: []int{4},
			wantDel: []int{3},
		},
	}
	for _, tC := range testCases {
		tC := tC

		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()

			added, deleted, err := slices.Diff(tC.x, tC.y)

			if !cmp.Equal(added, tC.wantAdd, cmpopts.EquateEmpty()) {
				t.Errorf("wantAdd = %v, got = %v", tC.wantAdd, added)
			}

			if !cmp.Equal(deleted, tC.wantDel, cmpopts.EquateEmpty()) {
				t.Errorf("wantDel = %v, got = %v", tC.wantDel, deleted)
			}

			if !cmp.Equal(err, tC.wantErr, cmpopts.EquateErrors()) {
				t.Errorf("wantErr = %v, got = %v", tC.wantErr, err)
			}
		})
	}
}

func BenchmarkStrDiff(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slices.Diff([]string{"a", "b", "c"}, []string{"a", "b", "d"})
	}
}

func BenchmarkStringerDiff(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slices.Diff([]ex{{"a"}, {"b"}, {"c"}}, []ex{{"a"}, {"b"}, {"d"}})
	}
}

func BenchmarkIntDiff(b *testing.B) {
	for n := 0; n < b.N; n++ {
		slices.Diff([]int{1, 2, 3}, []int{1, 2, 4})
	}
}
