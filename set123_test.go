//go:build go1.23
// +build go1.23

package mapset

import "testing"

func TestAll123(t *testing.T) {
	a := NewSet[string]()

	a.Add("Z")
	a.Add("Y")
	a.Add("X")
	a.Add("W")

	b := NewSet[string]()
	for elem := range Values(a) {
		b.Add(elem)
	}

	if !a.Equal(b) {
		t.Error("The sets are not equal after iterating (Each) through the first set")
	}

	var count int
	for range Values(a) {
		if count == 2 {
			break
		}
		count++
	}

	if count != 2 {
		t.Error("Iteration should stop on the way")
	}
}
