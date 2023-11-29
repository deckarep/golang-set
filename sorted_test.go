//go:build go1.21
// +build go1.21

/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 - 2022 Ralph Caraveo (deckarep@gmail.com)

Permission is hereby granted, free of charge, to any person obtaining a copy of
this software and associated documentation files (the "Software"), to deal in
the Software without restriction, including without limitation the rights to
use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
of the Software, and to permit persons to whom the Software is furnished to do
so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package mapset

import (
	"testing"
)

func Test_Sorted(t *testing.T) {
	test := func(t *testing.T, ctor func(vals ...string) Set[string]) {
		set := ctor("apple", "banana", "pear")
		sorted := Sorted(set)

		if len(sorted) != set.Cardinality() {
			t.Errorf("Length of slice is not the same as the set. Expected: %d. Actual: %d", set.Cardinality(), len(sorted))
		}

		if sorted[0] != "apple" {
			t.Errorf("Element 0 was not equal to apple: %s", sorted[0])
		}
		if sorted[1] != "banana" {
			t.Errorf("Element 1 was not equal to banana: %s", sorted[1])
		}
		if sorted[2] != "pear" {
			t.Errorf("Element 2 was not equal to pear: %s", sorted[2])
		}
	}

	t.Run("Safe", func(t *testing.T) {
		test(t, NewSet[string])
	})
	t.Run("Unsafe", func(t *testing.T) {
		test(t, NewThreadUnsafeSet[string])
	})
}
