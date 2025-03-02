//go:build go1.23
// +build go1.23

package mapset

import "testing"

func TestThreadUnsageSetElements(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2, 3)
	got := NewThreadUnsafeSet[int]()

	var gotPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				gotPanic = true
			}
		}()

		for i := range s.Elements {
			got.Add(i)
		}
	}()

	if gotPanic {
		t.Errorf("Elements should NOT panic if called under go1.23+")
	}

	if !s.Equal(got) {
		t.Errorf("Expected no difference, got: %v", s.Difference(got))
	}
}

func TestThreadsafeSetElements(t *testing.T) {
	s := NewSet(1, 2, 3)
	got := NewSet[int]()

	var gotPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				gotPanic = true
			}
		}()

		for i := range s.Elements {
			got.Add(i)
		}
	}()

	if gotPanic {
		t.Errorf("Elements should NOT panic if called under go1.23+")
	}

	if !s.Equal(got) {
		t.Errorf("Expected no difference, got: %v", s.Difference(got))
	}
}
