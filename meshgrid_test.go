package meshgrid

import (
	"strconv"
	"testing"
)

func TestSubTo(t *testing.T) {
	// Convert to [] and use a map to check all unique
	for _, dims := range [][]int{
		{2, 2},
		{3, 4},
		{2, 2, 2, 2},
		{4, 3, 2, 1},
		{5, 4, 3, 2},
		{2, 3, 4, 5},
	} {
		total := 1
		for _, v := range dims {
			total *= v
		}
		subs := make(map[string]struct{})
		for i := 0; i < total; i++ {
			sub := SubFor(nil, i, dims)
			if len(sub) != len(dims) {
				t.Errorf("Subscript has bad length")
			}
			sub2 := make([]int, len(sub))
			SubFor(sub2, i, dims)
			if len(sub) != len(sub2) {
				t.Errorf("Sub length mismatch when non-nil")
			}
			equal := true
			for j, v := range sub {
				if sub2[j] != v {
					equal = false
				}
			}
			if !equal {
				t.Errorf("Sub mismatch when non-nil")
			}

			// Turn the found subscript into a string to make sure all of the
			// generated strings are indential (so there isn't a bug where)
			// everything maps to 0 or something.
			var str string
			for _, v := range sub {
				str += strconv.Itoa(v) + "_"
			}
			_, ok := subs[str]
			if ok {
				t.Errorf("duplicate string %s", str)
			}
			subs[str] = struct{}{}

			idx := IdxFor(sub, dims)
			if idx != i {
				t.Errorf("Idx mismatch. Want %v, got %v", i, idx)
			}
		}
	}
}
