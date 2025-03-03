//go:build !go1.23
// +build !go1.23

package mapset

import "testing"

func TestThreadUnsafeSetElements(t *testing.T) {
	s := NewThreadUnsafeSet(1, 2, 3)

	var gotPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				gotPanic = true
			}
		}()

		s.Elements()(1)
	}()

	if !gotPanic {
		t.Errorf("Elements should panic if called under go1.22")
	}
}

func TestNewThreadUnsafeSetFromSeq(t *testing.T) {
	var gotPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				gotPanic = true
			}
		}()

		NewThreadUnsafeSetFromSeq[int](func(func(int) bool) {})
	}()

	if !gotPanic {
		t.Errorf("Elements should panic if called under go1.22")
	}
}

func TestNewThreadUnsafeSetFromSeq2Keys(t *testing.T) {
	var gotPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				gotPanic = true
			}
		}()

		NewThreadUnsafeSetFromSeq2Keys[int, string](func(func(int, string) bool) {})
	}()

	if !gotPanic {
		t.Errorf("Elements should panic if called under go1.22")
	}
}

func TestNewThreadUnsafeSetFromSeq2Values(t *testing.T) {
	var gotPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				gotPanic = true
			}
		}()

		NewThreadUnsafeSetFromSeq2Values[int, string](func(func(int, string) bool) {})
	}()

	if !gotPanic {
		t.Errorf("Elements should panic if called under go1.22")
	}
}

func TestThreadsafeSetElements(t *testing.T) {
	s := NewSet(1, 2, 3)

	var gotPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				gotPanic = true
			}
		}()

		s.Elements()(1)
	}()

	if !gotPanic {
		t.Errorf("Elements should panic if called under go1.22")
	}
}

func TestNewSetFromSeq(t *testing.T) {
	var gotPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				gotPanic = true
			}
		}()

		NewSetFromSeq[int](func(func(int) bool) {})

	}()

	if !gotPanic {
		t.Errorf("Elements should panic if called under go1.22")
	}
}

func TestNewSetFromSeq2Keys(t *testing.T) {
	var gotPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				gotPanic = true
			}
		}()

		NewSetFromSeq2Keys[int, string](func(func(int, string) bool) {})
	}()

	if !gotPanic {
		t.Errorf("Elements should panic if called under go1.22")
	}
}

func TestNewSetFromSeq2Values(t *testing.T) {
	var gotPanic bool
	func() {
		defer func() {
			if r := recover(); r != nil {
				gotPanic = true
			}
		}()

		NewSetFromSeq2Values[int, string](func(func(int, string) bool) {})
	}()

	if !gotPanic {
		t.Errorf("Elements should panic if called under go1.22")
	}
}
