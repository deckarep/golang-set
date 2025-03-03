//go:build go1.23
// +build go1.23

package mapset

import (
	"slices"
	"testing"
)

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

		for i := range s.Elements() {
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

func TestNewThreadUnsafeSetFromSeq(t *testing.T) {
	names := []string{"Alice", "Bob", "Vera", "Bob"}
	got := NewThreadUnsafeSetFromSeq(slices.Values(names))
	expect := NewThreadUnsafeSet(names...)
	if !got.Equal(expect) {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}

func TestNewThreadUnsafeSetFromSeq2Keys(t *testing.T) {
	names := []string{"Alice", "Bob", "Vera", "Bob"}
	got := NewThreadUnsafeSetFromSeq2Keys(slices.All(names))
	expect := NewThreadUnsafeSet(0, 1, 2, 3)
	if !got.Equal(expect) {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}

func TestNewThreadUnsafeSetFromSeqValues(t *testing.T) {
	names := []string{"Alice", "Bob", "Vera", "Bob"}
	got := NewThreadUnsafeSetFromSeq2Values(slices.All(names))
	expect := NewThreadUnsafeSet(names...)
	if !got.Equal(expect) {
		t.Errorf("Expected %v, got %v", expect, got)
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

		for i := range s.Elements() {
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

func TestNewSetFromSeq(t *testing.T) {
	names := []string{"Alice", "Bob", "Vera", "Bob"}
	got := NewSetFromSeq(slices.Values(names))
	expect := NewSet(names...)
	if !got.Equal(expect) {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}

func TestNewSetFromSeq2Keys(t *testing.T) {
	names := []string{"Alice", "Bob", "Vera", "Bob"}
	got := NewSetFromSeq2Keys(slices.All(names))
	expect := NewSet(0, 1, 2, 3)
	if !got.Equal(expect) {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}

func TestNewSetFromSeqValues(t *testing.T) {
	names := []string{"Alice", "Bob", "Vera", "Bob"}
	got := NewSetFromSeq2Values(slices.All(names))
	expect := NewSet(names...)
	if !got.Equal(expect) {
		t.Errorf("Expected %v, got %v", expect, got)
	}
}
