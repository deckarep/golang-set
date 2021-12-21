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
	"fmt"
	"testing"
)

// Want to test numbers, strings and some kind of struct.
// That will be good enough.

func Test_Generic(t *testing.T) {
	a := newThreadUnsafeSetGeneric[int]()
	if a.Cardinality() != 0 {
		t.Error("should be empty")
	}

	a.Add(5)
	if a.Cardinality() != 1 {
		t.Error("should have one element")
	}

	a.Add(4)
	if a.Cardinality() != 2 {
		t.Error("should have 2 elements")
	}

	s := a.ToSlice()
	fmt.Println(s)
}

func Test_Generic_String(t *testing.T) {
	a := newThreadUnsafeSetGeneric[int]()
	a.Add(5)
	a.Add(4)

	fmt.Println(a.String())
}

func Test_Powerset(t *testing.T) {
	a := newThreadUnsafeSetGeneric[string]()
	a.Add("x")
	a.Add("y")
	a.Add("z")

	ps := a.PowerSet()
	fmt.Println(ps)
}
