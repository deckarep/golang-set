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
