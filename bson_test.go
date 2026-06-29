//go:build !tinygo

/*
Open Source Initiative OSI - The MIT License (MIT):Licensing

The MIT License (MIT)
Copyright (c) 2013 - 2026 Ralph Caraveo (deckarep@gmail.com)

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

	"go.mongodb.org/mongo-driver/bson"
)

func Test_UnmarshalBSONValue(t *testing.T) {
	tp, s, initErr := bson.MarshalValue(
		bson.A{"1", "2", "3", "test"},
	)

	if initErr != nil {
		t.Errorf("Init Error should be nil: %v", initErr)

		return
	}

	if tp != bson.TypeArray {
		t.Errorf("Encoded Type should be bson.Array, got: %v", tp)

		return
	}

	expected := NewSet("1", "2", "3", "test")
	actual := NewSet[string]()
	err := bson.UnmarshalValue(bson.TypeArray, s, actual)
	if err != nil {
		t.Errorf("Error should be nil: %v", err)
	}

	if !expected.Equal(actual) {
		t.Errorf("Expected no difference, got: %v", expected.Difference(actual))
	}
}

func TestThreadUnsafeSet_UnmarshalBSONValue(t *testing.T) {
	tp, s, initErr := bson.MarshalValue(
		bson.A{int64(1), int64(2), int64(3)},
	)

	if initErr != nil {
		t.Errorf("Init Error should be nil: %v", initErr)

		return
	}

	if tp != bson.TypeArray {
		t.Errorf("Encoded Type should be bson.Array, got: %v", tp)

		return
	}

	expected := NewThreadUnsafeSet[int64](1, 2, 3)
	actual := NewThreadUnsafeSet[int64]()
	err := actual.UnmarshalBSONValue(bson.TypeArray, []byte(s))
	if err != nil {
		t.Errorf("Error should be nil: %v", err)
	}
	if !expected.Equal(actual) {
		t.Errorf("Expected no difference, got: %v", expected.Difference(actual))
	}
}

func Test_MarshalBSONValue(t *testing.T) {
	expected := NewSet("1", "test")

	_, b, err := bson.MarshalValue(
		NewSet("1", "test"),
	)
	if err != nil {
		t.Errorf("Error should be nil: %v", err)
	}

	actual := NewSet[string]()
	err = bson.UnmarshalValue(bson.TypeArray, b, actual)
	if err != nil {
		t.Errorf("Error should be nil: %v", err)
	}

	if !expected.Equal(actual) {
		t.Errorf("Expected no difference, got: %v", expected.Difference(actual))
	}
}
